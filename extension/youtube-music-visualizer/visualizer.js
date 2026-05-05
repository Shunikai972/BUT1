/*
 * YouTube Music 3D Visualizer
 * Single-file content script + page script.
 *
 * The content-script half injects this same file into the page's MAIN world.
 * The page-script half can then read the real YouTube Music media element and
 * feed a Web Audio analyser without fighting the extension isolated world.
 */
(function ytmvBootstrap() {
  "use strict";

  var PAGE_MARKER = "ytmv-page-script";
  var currentScript = document.currentScript;
  var isPageScript = !!(
    currentScript &&
    currentScript.dataset &&
    currentScript.dataset.ytmvRole === PAGE_MARKER
  );

  if (!isPageScript) {
    injectIntoPage();
    return;
  }

  startVisualizer();

  function injectIntoPage() {
    try {
      var runtime =
        (typeof browser !== "undefined" && browser.runtime) ||
        (typeof chrome !== "undefined" && chrome.runtime);

      if (!runtime || !runtime.getURL) {
        console.warn("[YTMV] Extension runtime unavailable; page injection skipped.");
        return;
      }

      var script = document.createElement("script");
      script.src = runtime.getURL("visualizer.js") + "?ytmv-main=" + Date.now();
      script.dataset.ytmvRole = PAGE_MARKER;
      script.async = false;
      script.onload = function () {
        script.remove();
      };
      script.onerror = function () {
        console.warn("[YTMV] Page injection failed. Check web_accessible_resources.");
        script.remove();
      };

      (document.head || document.documentElement).appendChild(script);
    } catch (error) {
      console.warn("[YTMV] Unable to inject page script:", error);
    }
  }

  function startVisualizer() {
    if (window.__YTMV_APP__ && typeof window.__YTMV_APP__.destroy === "function") {
      window.__YTMV_APP__.destroy();
    }

    var STORE_KEY = "ytmv.visualizer.settings.v3";
    var DEFAULTS = {
      enabled: true,
      collapsed: false,
      style: "vortex",
      sensitivity: 1.62,
      opacity: 0.94,
      density: 0.82,
      performance: false,
      panelX: null,
      panelY: null
    };

    var config = loadConfig();
    var cleanup = [];
    var rafId = 0;
    var host = null;
    var shadow = null;
    var canvas = null;
    var ctx = null;
    var panel = null;
    var observer = null;
    var cssWidth = 1;
    var cssHeight = 1;
    var pixelRatio = 1;
    var lastFrameAt = 0;
    var lastPanelAt = 0;
    var lastEnsureAt = 0;

    var audio = {
      context: null,
      analyser: null,
      frequency: null,
      waveform: null,
      source: null,
      output: null,
      stream: null,
      media: null,
      sourceKind: "none",
      connectedAt: 0,
      lastConnectAttemptAt: 0,
      lastSignalAt: 0,
      lastReconnectAt: 0,
      status: "searching",
      error: ""
    };

    var beat = {
      history: [],
      historySize: 58,
      index: 0,
      filled: false,
      lastAt: 0,
      intervals: [],
      bpm: 0,
      pulse: 0
    };

    var metrics = {
      bass: 0,
      lowMid: 0,
      mid: 0,
      high: 0,
      volume: 0,
      beat: false,
      bpm: 0,
      hasSignal: false
    };

    var smooth = {
      bass: 0,
      lowMid: 0,
      mid: 0,
      high: 0,
      volume: 0,
      pulse: 0
    };

    var scene = {
      particles: [],
      sparks: [],
      rings: [],
      phase: Math.random() * Math.PI * 2,
      hue: 184,
      quality: 1,
      reconnectNoticeUntil: 0
    };

    var elements = {};

    init();

    window.YTMVisualizer = {
      toggle: function () {
        setEnabled(!config.enabled);
      },
      setStyle: function (style) {
        if (["vortex", "sphere", "aurora", "spectrum"].indexOf(style) !== -1) {
          config.style = style;
          saveConfig();
        }
      },
      reconnect: function () {
        reconnectAudio("api");
      },
      getState: function () {
        return {
          enabled: config.enabled,
          style: config.style,
          source: audio.sourceKind,
          status: audio.status,
          signal: metrics.hasSignal,
          bpm: metrics.bpm
        };
      },
      destroy: destroy
    };
    window.__YTMV_APP__ = window.YTMVisualizer;

    function init() {
      createUI();
      resizeCanvas();
      createParticles();
      installEventHandlers();
      installMediaObserver();
      ensureAudio("boot");

      if (config.enabled) {
        startLoop();
      } else {
        applyEnabledState();
      }

      console.info("[YTMV] Visualizer ready. Looking for YouTube Music media.");
    }

    function loadConfig() {
      var loaded = {};
      try {
        loaded = JSON.parse(localStorage.getItem(STORE_KEY) || "{}");
      } catch (error) {
        loaded = {};
      }
      return assign({}, DEFAULTS, loaded);
    }

    function saveConfig() {
      try {
        localStorage.setItem(
          STORE_KEY,
          JSON.stringify({
            enabled: config.enabled,
            collapsed: config.collapsed,
            style: config.style,
            sensitivity: config.sensitivity,
            opacity: config.opacity,
            density: config.density,
            performance: config.performance,
            panelX: config.panelX,
            panelY: config.panelY
          })
        );
      } catch (error) {
        // localStorage may be blocked; the extension still works without it.
      }
    }

    function createUI() {
      host = document.createElement("div");
      host.id = "ytmv-extension-root";
      host.style.cssText = [
        "position:fixed",
        "inset:0",
        "z-index:2147483000",
        "pointer-events:none",
        "contain:layout style size",
        "font-size:14px"
      ].join(";");

      shadow = host.attachShadow({ mode: "open" });

      var style = document.createElement("style");
      style.textContent = getStyles();

      canvas = document.createElement("canvas");
      canvas.className = "visualizer";

      panel = document.createElement("section");
      panel.className = "panel";
      buildPanel(panel);

      shadow.appendChild(style);
      shadow.appendChild(canvas);
      shadow.appendChild(panel);
      document.documentElement.appendChild(host);

      elements.power = shadow.querySelector("[data-power]");
      elements.collapse = shadow.querySelector("[data-collapse]");
      elements.style = shadow.querySelector("[data-style]");
      elements.sensitivity = shadow.querySelector("[data-sensitivity]");
      elements.opacity = shadow.querySelector("[data-opacity]");
      elements.density = shadow.querySelector("[data-density]");
      elements.performance = shadow.querySelector("[data-performance]");
      elements.reconnect = shadow.querySelector("[data-reconnect]");
      elements.status = shadow.querySelector("[data-status]");
      elements.source = shadow.querySelector("[data-source]");
      elements.bass = shadow.querySelector("[data-bass]");
      elements.volume = shadow.querySelector("[data-volume]");
      elements.bpm = shadow.querySelector("[data-bpm]");

      elements.style.value = config.style;
      elements.sensitivity.value = String(config.sensitivity);
      elements.opacity.value = String(config.opacity);
      elements.density.value = String(config.density);
      elements.performance.checked = !!config.performance;

      if (config.panelX !== null && config.panelY !== null) {
        panel.style.left = config.panelX + "px";
        panel.style.top = config.panelY + "px";
        panel.style.right = "auto";
        panel.style.bottom = "auto";
      }

      applyEnabledState();
      applyCollapsedState();
    }

    function buildPanel(root) {
      var topbar = makeNode("div", "topbar");
      topbar.setAttribute("data-drag-handle", "true");

      var brand = makeNode("div", "brand");
      brand.appendChild(makeNode("span", "signal-dot"));
      brand.appendChild(makeNode("span", "title", "YTMV"));

      var source = makeNode("span", "source", "media");
      source.setAttribute("data-source", "");
      brand.appendChild(source);

      var topActions = makeNode("div", "top-actions");
      topActions.appendChild(makeButton("mini-button", "ON", "data-power", "Activer / desactiver"));
      topActions.appendChild(makeButton("mini-button", "-", "data-collapse", "Reduire le menu"));

      topbar.appendChild(brand);
      topbar.appendChild(topActions);

      var body = makeNode("div", "body");

      var styleSelect = document.createElement("select");
      styleSelect.setAttribute("data-style", "");
      addOption(styleSelect, "vortex", "Vortex");
      addOption(styleSelect, "sphere", "Sphere");
      addOption(styleSelect, "aurora", "Aurora");
      addOption(styleSelect, "spectrum", "Spectrum");
      body.appendChild(makeField("Style", styleSelect));

      body.appendChild(makeField("Gain", makeRange("data-sensitivity", "0.65", "2.6", "0.05")));
      body.appendChild(makeField("Opacite", makeRange("data-opacity", "0.18", "1", "0.02")));
      body.appendChild(makeField("Densite", makeRange("data-density", "0.45", "1.45", "0.05")));

      var performance = makeNode("label", "toggle-row");
      var performanceInput = document.createElement("input");
      performanceInput.type = "checkbox";
      performanceInput.setAttribute("data-performance", "");
      performance.appendChild(performanceInput);
      performance.appendChild(makeNode("span", "", "Mode perf"));
      body.appendChild(performance);

      var meters = makeNode("div", "meters");
      meters.appendChild(makeMeter("Bass", "data-bass"));
      meters.appendChild(makeMeter("Vol", "data-volume"));

      var bpm = makeNode("div", "bpm");
      bpm.appendChild(makeNode("span", "", "BPM"));
      var bpmValue = makeNode("strong", "", "--");
      bpmValue.setAttribute("data-bpm", "");
      bpm.appendChild(bpmValue);
      meters.appendChild(bpm);
      body.appendChild(meters);

      body.appendChild(makeButton("reconnect", "Reconnecter audio", "data-reconnect", ""));

      var status = makeNode("div", "status", "Recherche du lecteur...");
      status.setAttribute("data-status", "");
      body.appendChild(status);

      root.appendChild(topbar);
      root.appendChild(body);
    }

    function makeNode(tagName, className, text) {
      var node = document.createElement(tagName);
      if (className) {
        node.className = className;
      }
      if (typeof text === "string") {
        node.textContent = text;
      }
      return node;
    }

    function makeButton(className, text, dataName, title) {
      var button = makeNode("button", className, text);
      button.type = "button";
      button.setAttribute(dataName, "");
      if (title) {
        button.title = title;
      }
      return button;
    }

    function makeField(labelText, control) {
      var label = makeNode("label", "field");
      label.appendChild(makeNode("span", "", labelText));
      label.appendChild(control);
      return label;
    }

    function makeRange(dataName, min, max, step) {
      var input = document.createElement("input");
      input.type = "range";
      input.min = min;
      input.max = max;
      input.step = step;
      input.setAttribute(dataName, "");
      return input;
    }

    function addOption(select, value, label) {
      var option = document.createElement("option");
      option.value = value;
      option.textContent = label;
      select.appendChild(option);
    }

    function makeMeter(labelText, dataName) {
      var meter = makeNode("div", "meter");
      meter.appendChild(makeNode("span", "", labelText));
      var bar = document.createElement("i");
      bar.setAttribute(dataName, "");
      meter.appendChild(bar);
      return meter;
    }

    function getStyles() {
      return [
        ":host{all:initial}",
        ".visualizer{position:fixed;inset:0;width:100%;height:100%;pointer-events:none;opacity:var(--ytmv-opacity,.94);mix-blend-mode:screen;filter:saturate(1.35) contrast(1.12)}",
        ".panel{position:fixed;right:18px;bottom:18px;width:286px;box-sizing:border-box;pointer-events:auto;color:#f7fbff;font-family:Inter,ui-sans-serif,system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;background:linear-gradient(145deg,rgba(11,15,22,.84),rgba(20,27,39,.72));border:1px solid rgba(141,225,255,.32);border-radius:8px;box-shadow:0 22px 70px rgba(0,0,0,.46),0 0 34px rgba(72,205,255,.18);backdrop-filter:blur(18px) saturate(135%);overflow:hidden;user-select:none}",
        ".topbar{height:42px;display:flex;align-items:center;justify-content:space-between;padding:0 8px 0 12px;box-sizing:border-box;cursor:grab;border-bottom:1px solid rgba(255,255,255,.08);background:linear-gradient(90deg,rgba(73,196,255,.14),rgba(255,82,178,.10),rgba(255,190,82,.08))}",
        ".topbar:active{cursor:grabbing}",
        ".brand{display:flex;align-items:center;gap:8px;min-width:0}",
        ".signal-dot{width:9px;height:9px;border-radius:999px;background:#ffc857;box-shadow:0 0 12px rgba(255,200,87,.75)}",
        ".panel.has-signal .signal-dot{background:#45f0a0;box-shadow:0 0 14px rgba(69,240,160,.9)}",
        ".panel.no-media .signal-dot{background:#ff5d73;box-shadow:0 0 14px rgba(255,93,115,.8)}",
        ".title{font-weight:760;font-size:12px;line-height:1;letter-spacing:0;color:#fff}",
        ".source{max-width:118px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;color:#acd4e8;font-size:11px;line-height:1}",
        ".top-actions{display:flex;gap:6px}",
        "button,select,input{font:inherit}",
        ".mini-button{min-width:34px;height:26px;border:1px solid rgba(255,255,255,.14);border-radius:6px;background:rgba(255,255,255,.08);color:#f7fbff;cursor:pointer}",
        ".mini-button:hover,.reconnect:hover{background:rgba(255,255,255,.14)}",
        ".body{display:grid;gap:10px;padding:12px}",
        ".field{display:grid;grid-template-columns:64px 1fr;align-items:center;gap:10px;color:#cfe0eb;font-size:12px}",
        ".field span,.toggle-row span{color:#cfe0eb;font-size:12px}",
        "select{width:100%;height:30px;color:#f8fcff;background:rgba(5,8,13,.62);border:1px solid rgba(141,225,255,.28);border-radius:6px;padding:0 8px;outline:none}",
        "input[type='range']{width:100%;accent-color:#55d9ff}",
        ".toggle-row{display:flex;align-items:center;gap:9px;color:#cfe0eb;font-size:12px}",
        ".toggle-row input{accent-color:#55d9ff}",
        ".meters{display:grid;grid-template-columns:1fr 1fr 62px;gap:8px;align-items:stretch}",
        ".meter{height:36px;border-radius:6px;background:rgba(255,255,255,.06);border:1px solid rgba(255,255,255,.08);position:relative;overflow:hidden}",
        ".meter span{position:absolute;left:8px;top:6px;z-index:1;color:#dcecf5;font-size:10px}",
        ".meter i{position:absolute;left:0;bottom:0;height:4px;width:0%;background:linear-gradient(90deg,#55d9ff,#ff5bb8,#ffd166);box-shadow:0 0 14px rgba(85,217,255,.55)}",
        ".bpm{height:36px;border-radius:6px;background:rgba(255,255,255,.06);border:1px solid rgba(255,255,255,.08);display:grid;place-items:center}",
        ".bpm span{font-size:9px;color:#adc7d7;line-height:1}",
        ".bpm strong{font-size:14px;color:#ffffff;line-height:1}",
        ".reconnect{height:32px;border:1px solid rgba(141,225,255,.28);border-radius:6px;background:rgba(85,217,255,.08);color:#f7fbff;cursor:pointer}",
        ".status{min-height:16px;color:#adc7d7;font-size:11px;line-height:1.35}",
        ".panel.is-collapsed{width:168px}",
        ".panel.is-collapsed .body{display:none}",
        ".panel.is-disabled .visual-state{opacity:.45}",
        "@media (max-width:520px){.panel{right:10px;bottom:10px;width:min(286px,calc(100vw - 20px))}.panel.is-collapsed{width:154px}.source{max-width:82px}}"
      ].join("\n");
    }

    function installEventHandlers() {
      on(window, "resize", function () {
        resizeCanvas();
        boundPanelToViewport();
      });

      on(document, "visibilitychange", function () {
        if (document.hidden) {
          stopLoop();
        } else if (config.enabled) {
          ensureAudio("visible");
          startLoop();
        }
      });

      on(document, "play", function (event) {
        if (isMediaElement(event.target)) {
          setTimeout(function () {
            ensureAudio("play");
          }, 80);
        }
      }, true);

      on(document, "pause", function (event) {
        if (isMediaElement(event.target) && event.target === audio.media) {
          forceSilence("pause");
          clearCanvas();
          updatePanel(performance.now(), true);
        }
      }, true);

      on(document, "ended", function (event) {
        if (isMediaElement(event.target) && event.target === audio.media) {
          forceSilence("ended");
          clearCanvas();
          updatePanel(performance.now(), true);
        }
      }, true);

      ["pointerdown", "keydown", "touchstart"].forEach(function (name) {
        on(window, name, function () {
          unlockAudio(name);
        }, { capture: true, passive: true });
      });

      elements.power.addEventListener("click", function () {
        unlockAudio("power");
        setEnabled(!config.enabled);
      });

      elements.collapse.addEventListener("click", function () {
        config.collapsed = !config.collapsed;
        applyCollapsedState();
        saveConfig();
      });

      elements.style.addEventListener("change", function () {
        config.style = elements.style.value;
        saveConfig();
      });

      elements.sensitivity.addEventListener("input", function () {
        config.sensitivity = parseFloat(elements.sensitivity.value);
        saveConfig();
      });

      elements.opacity.addEventListener("input", function () {
        config.opacity = parseFloat(elements.opacity.value);
        applyOpacity();
        saveConfig();
      });

      elements.density.addEventListener("input", function () {
        config.density = parseFloat(elements.density.value);
        createParticles();
        saveConfig();
      });

      elements.performance.addEventListener("change", function () {
        config.performance = !!elements.performance.checked;
        configureAnalyser();
        createParticles();
        saveConfig();
      });

      elements.reconnect.addEventListener("click", function () {
        unlockAudio("reconnect");
        reconnectAudio("button");
      });

      makePanelDraggable();
    }

    function installMediaObserver() {
      observer = new MutationObserver(function () {
        if (!audio.media || !document.documentElement.contains(audio.media)) {
          ensureAudio("mutation");
        }
      });
      observer.observe(document.documentElement, { childList: true, subtree: true });
      cleanup.push(function () {
        observer.disconnect();
      });
    }

    function on(target, type, handler, options) {
      target.addEventListener(type, handler, options);
      cleanup.push(function () {
        target.removeEventListener(type, handler, options);
      });
    }

    function resizeCanvas() {
      cssWidth = Math.max(1, window.innerWidth || document.documentElement.clientWidth || 1);
      cssHeight = Math.max(1, window.innerHeight || document.documentElement.clientHeight || 1);
      pixelRatio = config.performance ? 0.86 : Math.min(1.15, window.devicePixelRatio || 1);
      canvas.width = Math.max(1, Math.floor(cssWidth * pixelRatio));
      canvas.height = Math.max(1, Math.floor(cssHeight * pixelRatio));
      canvas.style.setProperty("--ytmv-opacity", String(config.opacity));
      ctx = canvas.getContext("2d", { alpha: true, desynchronized: true });
      ctx.setTransform(pixelRatio, 0, 0, pixelRatio, 0, 0);
      ctx.lineCap = "round";
      ctx.lineJoin = "round";
      ctx.imageSmoothingEnabled = !config.performance;
      createParticles();
    }

    function applyOpacity() {
      if (canvas) {
        canvas.style.setProperty("--ytmv-opacity", String(config.opacity));
      }
    }

    function setEnabled(enabled) {
      config.enabled = !!enabled;
      applyEnabledState();
      saveConfig();

      if (config.enabled) {
        unlockAudio("enable");
        startLoop();
      } else {
        stopLoop();
        clearCanvas();
      }
    }

    function applyEnabledState() {
      if (!canvas || !panel) {
        return;
      }
      canvas.style.display = config.enabled ? "block" : "none";
      panel.classList.toggle("is-disabled", !config.enabled);
      elements.power.textContent = config.enabled ? "ON" : "OFF";
    }

    function applyCollapsedState() {
      panel.classList.toggle("is-collapsed", !!config.collapsed);
      elements.collapse.textContent = config.collapsed ? "+" : "-";
      elements.collapse.title = config.collapsed ? "Agrandir le menu" : "Reduire le menu";
      boundPanelToViewport();
    }

    function makePanelDraggable() {
      var handle = shadow.querySelector("[data-drag-handle]");
      var drag = null;

      handle.addEventListener("pointerdown", function (event) {
        if (event.target.closest("button")) {
          return;
        }

        var rect = panel.getBoundingClientRect();
        drag = {
          x: event.clientX - rect.left,
          y: event.clientY - rect.top
        };
        panel.style.left = rect.left + "px";
        panel.style.top = rect.top + "px";
        panel.style.right = "auto";
        panel.style.bottom = "auto";
        handle.setPointerCapture(event.pointerId);
      });

      handle.addEventListener("pointermove", function (event) {
        if (!drag) {
          return;
        }

        var nextX = clamp(event.clientX - drag.x, 8, window.innerWidth - panel.offsetWidth - 8);
        var nextY = clamp(event.clientY - drag.y, 8, window.innerHeight - panel.offsetHeight - 8);
        panel.style.left = nextX + "px";
        panel.style.top = nextY + "px";
        config.panelX = Math.round(nextX);
        config.panelY = Math.round(nextY);
      });

      handle.addEventListener("pointerup", function (event) {
        if (!drag) {
          return;
        }
        drag = null;
        try {
          handle.releasePointerCapture(event.pointerId);
        } catch (error) {
          // The pointer may already have been released.
        }
        saveConfig();
      });
    }

    function boundPanelToViewport() {
      if (!panel || config.panelX === null || config.panelY === null) {
        return;
      }

      var nextX = clamp(config.panelX, 8, window.innerWidth - panel.offsetWidth - 8);
      var nextY = clamp(config.panelY, 8, window.innerHeight - panel.offsetHeight - 8);
      panel.style.left = nextX + "px";
      panel.style.top = nextY + "px";
      config.panelX = Math.round(nextX);
      config.panelY = Math.round(nextY);
    }

    function ensureAudio(reason) {
      if (!initAudioContext()) {
        return false;
      }

      var media = findBestMediaElement();
      if (!media) {
        audio.status = "no-media";
        audio.error = "Lecteur YouTube Music introuvable";
        return false;
      }

      if (audio.media !== media || !audio.source) {
        var now = performance.now();
        if (audio.media === media && !audio.source && now - audio.lastConnectAttemptAt < 900) {
          return false;
        }
        connectMedia(media, reason);
      }

      return !!audio.source;
    }

    function initAudioContext() {
      if (audio.context) {
        return true;
      }

      var AudioContextCtor = window.AudioContext || window.webkitAudioContext;
      if (!AudioContextCtor) {
        audio.status = "unsupported";
        audio.error = "Web Audio API indisponible";
        return false;
      }

      try {
        audio.context = new AudioContextCtor({ latencyHint: "interactive" });
      } catch (error) {
        try {
          audio.context = new AudioContextCtor();
        } catch (innerError) {
          audio.status = "unsupported";
          audio.error = "AudioContext refuse par le navigateur";
          return false;
        }
      }

      audio.analyser = audio.context.createAnalyser();
      configureAnalyser();
      audio.lastSignalAt = performance.now();
      return true;
    }

    function configureAnalyser() {
      if (!audio.analyser) {
        return;
      }
      audio.analyser.fftSize = config.performance ? 1024 : 2048;
      audio.analyser.smoothingTimeConstant = config.performance ? 0.72 : 0.8;
      audio.frequency = new Uint8Array(audio.analyser.frequencyBinCount);
      audio.waveform = new Uint8Array(audio.analyser.fftSize);
    }

    function unlockAudio(reason) {
      if (!ensureAudio(reason)) {
        return;
      }

      if (audio.context && audio.context.state === "suspended") {
        audio.context.resume().catch(function (error) {
          audio.error = "Clique encore une fois pour deverrouiller l'audio";
          console.warn("[YTMV] AudioContext resume failed:", error);
        });
      }
    }

    function reconnectAudio(reason) {
      audio.lastReconnectAt = performance.now();
      scene.reconnectNoticeUntil = audio.lastReconnectAt + 1300;
      disconnectCurrentSource();
      audio.media = null;
      audio.sourceKind = "none";
      ensureAudio(reason);
    }

    function findBestMediaElement() {
      var candidates = Array.prototype.slice.call(document.querySelectorAll("video,audio"));
      var best = null;
      var bestScore = -1;

      candidates.forEach(function (media) {
        if (!isMediaElement(media)) {
          return;
        }

        var rect = safeRect(media);
        var score = 0;
        var src = media.currentSrc || media.src || "";

        if (media.tagName.toLowerCase() === "video") {
          score += 20;
        }
        if (media.classList && media.classList.contains("html5-main-video")) {
          score += 40;
        }
        if (media.closest && media.closest("#movie_player")) {
          score += 35;
        }
        if (!media.paused) {
          score += 45;
        }
        if (!media.ended) {
          score += 8;
        }
        if (media.readyState >= 2) {
          score += media.readyState * 8;
        }
        if (media.currentTime > 0) {
          score += 12;
        }
        if (src) {
          score += 8;
        }
        if (rect.width > 0 || rect.height > 0) {
          score += 4;
        }
        if (media.muted || media.volume === 0) {
          score -= 10;
        }

        if (score > bestScore) {
          best = media;
          bestScore = score;
        }
      });

      return best;
    }

    function isMediaElement(value) {
      return value instanceof HTMLMediaElement;
    }

    function safeRect(element) {
      try {
        return element.getBoundingClientRect();
      } catch (error) {
        return { width: 0, height: 0 };
      }
    }

    function connectMedia(media, reason) {
      if (!audio.context || !audio.analyser) {
        return false;
      }

      disconnectCurrentSource();
      audio.lastConnectAttemptAt = performance.now();
      audio.media = media;
      audio.error = "";

      var connected = connectWithCaptureStream(media);
      if (!connected) {
        connected = connectWithMediaElementSource(media);
      }

      if (connected) {
        audio.connectedAt = performance.now();
        audio.lastSignalAt = performance.now();
        audio.status = "connected";
        attachMediaListeners(media);
        console.info("[YTMV] Audio connected via " + audio.sourceKind + " (" + reason + ").");
      } else {
        audio.status = "blocked";
        audio.error = "Flux audio inaccessible. Essaie Reconnecter apres lecture.";
        console.warn("[YTMV] Unable to connect to media element.");
      }

      return connected;
    }

    function connectWithCaptureStream(media) {
      var capture = media.captureStream || media.mozCaptureStream;
      if (!capture) {
        return false;
      }

      try {
        var stream = capture.call(media);
        if (!stream || !stream.getAudioTracks || stream.getAudioTracks().length === 0) {
          return false;
        }

        audio.stream = stream;
        audio.source = audio.context.createMediaStreamSource(stream);
        audio.source.connect(audio.analyser);
        audio.sourceKind = "captureStream";
        return true;
      } catch (error) {
        audio.error = error.message || "captureStream refuse";
        return false;
      }
    }

    function connectWithMediaElementSource(media) {
      if (!canUseMediaElementSource(media)) {
        return false;
      }

      try {
        var source = media.__ytmvMediaElementSource;
        if (!source) {
          source = audio.context.createMediaElementSource(media);
          Object.defineProperty(media, "__ytmvMediaElementSource", {
            value: source,
            configurable: true
          });
        }

        audio.output = audio.context.createGain();
        audio.output.gain.value = 1;
        source.connect(audio.analyser);
        source.connect(audio.output);
        audio.output.connect(audio.context.destination);
        audio.source = source;
        audio.sourceKind = "mediaElement";
        return true;
      } catch (error) {
        audio.error = error.message || "MediaElementAudioSource refuse";
        return false;
      }
    }

    function canUseMediaElementSource(media) {
      var src = media.currentSrc || media.src || "";
      if (!src) {
        return true;
      }
      return (
        src.indexOf("blob:") === 0 ||
        src.indexOf("data:") === 0 ||
        src.indexOf(window.location.origin) === 0
      );
    }

    function attachMediaListeners(media) {
      var reconnect = function () {
        setTimeout(function () {
          if (audio.media === media) {
            reconnectAudio("media-event");
          }
        }, 120);
      };

      ["loadedmetadata", "canplay", "emptied", "durationchange"].forEach(function (eventName) {
        media.addEventListener(eventName, reconnect, { once: true });
      });
    }

    function disconnectCurrentSource() {
      try {
        if (audio.source) {
          audio.source.disconnect();
        }
      } catch (error) {
        // Disconnection is best-effort; duplicated connects are harmlessly retried.
      }

      try {
        if (audio.output) {
          audio.output.disconnect();
        }
      } catch (error) {
        // Same best-effort cleanup.
      }

      audio.source = null;
      audio.output = null;
      audio.stream = null;
      audio.sourceKind = "none";
    }

    function readAudioMetrics(now) {
      metrics.beat = false;

      if (!audio.analyser || !audio.frequency || !audio.waveform) {
        return metrics;
      }

      if (isPlaybackStopped()) {
        forceSilence("paused");
        return metrics;
      }

      audio.analyser.getByteFrequencyData(audio.frequency);
      audio.analyser.getByteTimeDomainData(audio.waveform);

      var rawBass = frequencyBand(24, 150);
      var rawLowMid = frequencyBand(150, 520);
      var rawMid = frequencyBand(520, 2400);
      var rawHigh = frequencyBand(2400, 9800);
      var rawVolume = waveformRms();
      var gain = config.sensitivity;

      metrics.bass = clamp(Math.pow(rawBass * gain, 0.72), 0, 1);
      metrics.lowMid = clamp(Math.pow(rawLowMid * gain, 0.78), 0, 1);
      metrics.mid = clamp(Math.pow(rawMid * gain, 0.82), 0, 1);
      metrics.high = clamp(Math.pow(rawHigh * gain, 0.9), 0, 1);
      metrics.volume = clamp(rawVolume * 3.2 * gain, 0, 1);
      metrics.hasSignal = rawVolume > 0.006 || rawBass + rawLowMid + rawMid + rawHigh > 0.018;

      if (metrics.hasSignal) {
        audio.lastSignalAt = now;
      }

      detectBeat(now, metrics.bass, metrics.volume);
      metrics.beat = beat.pulse > 0.92;
      metrics.bpm = beat.bpm;

      smooth.bass += (metrics.bass - smooth.bass) * 0.18;
      smooth.lowMid += (metrics.lowMid - smooth.lowMid) * 0.16;
      smooth.mid += (metrics.mid - smooth.mid) * 0.14;
      smooth.high += (metrics.high - smooth.high) * 0.12;
      smooth.volume += (metrics.volume - smooth.volume) * 0.16;
      smooth.pulse = Math.max(beat.pulse, smooth.pulse * 0.9);

      if (
        audio.media &&
        !audio.media.paused &&
        now - audio.lastSignalAt > 2600 &&
        now - audio.lastReconnectAt > 3200
      ) {
        reconnectAudio("silence");
      }

      return metrics;
    }

    function isPlaybackStopped() {
      return !!(
        audio.media &&
        (audio.media.paused || audio.media.ended || audio.media.readyState < 2)
      );
    }

    function forceSilence(reason) {
      metrics.bass = 0;
      metrics.lowMid = 0;
      metrics.mid = 0;
      metrics.high = 0;
      metrics.volume = 0;
      metrics.beat = false;
      metrics.bpm = 0;
      metrics.hasSignal = false;

      smooth.bass = 0;
      smooth.lowMid = 0;
      smooth.mid = 0;
      smooth.high = 0;
      smooth.volume = 0;
      smooth.pulse = 0;

      beat.pulse = 0;
      beat.lastAt = 0;
      beat.bpm = 0;
      beat.history = [];
      beat.index = 0;
      beat.filled = false;
      beat.intervals = [];
      scene.sparks = [];
      scene.rings = [];
      audio.status = reason || "paused";
    }

    function frequencyBand(minHz, maxHz) {
      if (!audio.context || !audio.frequency || audio.frequency.length === 0) {
        return 0;
      }

      var nyquist = audio.context.sampleRate / 2;
      var start = clamp(Math.floor((minHz / nyquist) * audio.frequency.length), 0, audio.frequency.length - 1);
      var end = clamp(Math.ceil((maxHz / nyquist) * audio.frequency.length), start + 1, audio.frequency.length);
      var sum = 0;

      for (var i = start; i < end; i += 1) {
        sum += audio.frequency[i];
      }

      return sum / ((end - start) * 255);
    }

    function waveformRms() {
      var sum = 0;
      var length = audio.waveform.length;
      for (var i = 0; i < length; i += 2) {
        var sample = (audio.waveform[i] - 128) / 128;
        sum += sample * sample;
      }
      return Math.sqrt(sum / Math.max(1, length / 2));
    }

    function detectBeat(now, bass, volume) {
      var energy = bass * 0.78 + volume * 0.22;
      var history = beat.history;

      if (history.length < beat.historySize) {
        history.push(energy);
      } else {
        history[beat.index] = energy;
        beat.index = (beat.index + 1) % beat.historySize;
        beat.filled = true;
      }

      var count = history.length;
      var avg = 0;
      for (var i = 0; i < count; i += 1) {
        avg += history[i];
      }
      avg /= Math.max(1, count);

      var variance = 0;
      for (var j = 0; j < count; j += 1) {
        var diff = history[j] - avg;
        variance += diff * diff;
      }
      var std = Math.sqrt(variance / Math.max(1, count));
      var threshold = Math.max(0.085, avg + std * 1.28 + 0.018);
      var enoughHistory = beat.filled || history.length > 24;
      var isBeat = enoughHistory && energy > threshold && now - beat.lastAt > 245;

      if (isBeat) {
        if (beat.lastAt > 0) {
          var interval = now - beat.lastAt;
          if (interval >= 300 && interval <= 1500) {
            beat.intervals.push(interval);
            if (beat.intervals.length > 10) {
              beat.intervals.shift();
            }
            beat.bpm = estimateBpm(beat.intervals);
          }
        }
        beat.lastAt = now;
        beat.pulse = 1;
        spawnSparks();
      } else {
        beat.pulse *= 0.88;
      }
    }

    function estimateBpm(intervals) {
      if (!intervals.length) {
        return 0;
      }

      var sorted = intervals.slice().sort(function (a, b) {
        return a - b;
      });
      var trimmed = sorted.slice(Math.floor(sorted.length * 0.15), Math.ceil(sorted.length * 0.85));
      var total = 0;
      trimmed.forEach(function (value) {
        total += value;
      });

      var bpm = 60000 / (total / Math.max(1, trimmed.length));
      while (bpm < 72) {
        bpm *= 2;
      }
      while (bpm > 178) {
        bpm /= 2;
      }
      return Math.round(bpm);
    }

    function startLoop() {
      if (rafId || document.hidden) {
        return;
      }
      lastFrameAt = performance.now();
      rafId = requestAnimationFrame(frame);
    }

    function stopLoop() {
      if (rafId) {
        cancelAnimationFrame(rafId);
        rafId = 0;
      }
    }

    function frame(now) {
      rafId = 0;
      if (!config.enabled || document.hidden) {
        return;
      }

      var rawDt = Math.max(8, now - lastFrameAt);
      var dt = Math.min(48, rawDt);
      lastFrameAt = now;
      adjustQuality(rawDt);
      if (!audio.source || now - lastEnsureAt > 650) {
        lastEnsureAt = now;
        ensureAudio("frame");
      }
      readAudioMetrics(now);
      render(now, dt);
      updatePanel(now);
      rafId = requestAnimationFrame(frame);
    }

    function adjustQuality(dt) {
      if (config.performance) {
        scene.quality += (0.62 - scene.quality) * 0.12;
        return;
      }

      if (dt > 27) {
        scene.quality = Math.max(0.56, scene.quality - 0.08);
      } else if (dt < 18) {
        scene.quality = Math.min(1, scene.quality + 0.018);
      }
    }

    function render(now, dt) {
      clearCanvas();

      if (isPlaybackStopped()) {
        return;
      }

      var t = now * 0.001;
      scene.phase += dt * 0.00038 * (1 + smooth.volume * 1.8);
      scene.hue = (scene.hue + dt * 0.006 + smooth.high * 0.22) % 360;

      ctx.save();
      ctx.globalCompositeOperation = "lighter";

      drawEdgeGlow(t);
      drawWaveformRibbon(t);

      if (config.style === "vortex") {
        drawVortex(t, dt);
      } else if (config.style === "sphere") {
        drawSphere(t);
      } else if (config.style === "aurora") {
        drawAurora(t);
      } else {
        drawSpectrum(t);
      }

      drawBeatRings(dt);
      drawSparks(dt);
      ctx.restore();
    }

    function clearCanvas() {
      if (!ctx) {
        return;
      }
      ctx.clearRect(0, 0, cssWidth, cssHeight);
    }

    function drawEdgeGlow(t) {
      var intensity = 0.28 + smooth.volume * 0.54 + smooth.pulse * 0.4;
      var radius = Math.max(cssWidth, cssHeight) * 0.72;
      var gradient = ctx.createRadialGradient(cssWidth * 0.5, cssHeight * 0.44, 0, cssWidth * 0.5, cssHeight * 0.44, radius);
      gradient.addColorStop(0, hsla(scene.hue, 95, 62, 0.13 * intensity));
      gradient.addColorStop(0.42, hsla(scene.hue + 70, 92, 58, 0.11 * intensity));
      gradient.addColorStop(1, "rgba(0,0,0,0)");
      ctx.fillStyle = gradient;
      ctx.fillRect(0, 0, cssWidth, cssHeight);
    }

    function drawWaveformRibbon(t) {
      if (!audio.waveform || !metrics.hasSignal) {
        return;
      }

      var steps = config.performance ? 92 : Math.round(128 * scene.quality);
      var centerY = cssHeight * (0.58 + Math.sin(t * 0.45) * 0.018);
      var amp = cssHeight * (0.1 + smooth.bass * 0.22 + smooth.pulse * 0.08);
      var hue = scene.hue + 32 + smooth.high * 80;
      var gradient = ctx.createLinearGradient(0, centerY - amp, cssWidth, centerY + amp);
      gradient.addColorStop(0, hsla(hue + 120, 98, 66, 0.04));
      gradient.addColorStop(0.28, hsla(hue, 100, 68, 0.52));
      gradient.addColorStop(0.58, hsla(hue + 58, 100, 66, 0.7));
      gradient.addColorStop(1, hsla(hue + 180, 98, 62, 0.04));

      ctx.save();
      ctx.strokeStyle = gradient;
      ctx.lineWidth = 4 + smooth.volume * 8 + smooth.pulse * 5;
      ctx.shadowBlur = 0;
      ctx.beginPath();

      for (var i = 0; i <= steps; i += 1) {
        var ratio = i / steps;
        var waveIndex = Math.floor(ratio * (audio.waveform.length - 1));
        var wave = (audio.waveform[waveIndex] - 128) / 128;
        var freq = sampleFrequency(Math.pow(ratio, 1.12));
        var y = centerY + wave * amp * (0.55 + freq * 1.9) + Math.sin(ratio * Math.PI * 8 + t * 1.6) * amp * 0.08 * smooth.mid;
        var x = ratio * cssWidth;
        if (i === 0) {
          ctx.moveTo(x, y);
        } else {
          ctx.lineTo(x, y);
        }
      }
      ctx.stroke();

      ctx.globalAlpha = 0.32 + smooth.volume * 0.24;
      ctx.lineWidth *= 2.8;
      ctx.stroke();
      ctx.restore();
    }

    function drawVortex(t, dt) {
      var cx = cssWidth * 0.5;
      var cy = cssHeight * 0.43;
      var minSide = Math.min(cssWidth, cssHeight);
      var pulse = 1 + smooth.bass * 0.34 + smooth.pulse * 0.16;
      var inner = minSide * (0.15 + smooth.volume * 0.035);
      var outer = minSide * 0.43 * pulse;
      var count = Math.round((config.performance ? 64 : 92) * scene.quality);

      drawRadialBars(cx, cy, inner, outer, t, count, 0.58);
      drawParticleField(cx, cy, t, dt, 0.62, true);
      drawCore(cx, cy, minSide * 0.055 * pulse, t);
    }

    function drawSphere(t) {
      var cx = cssWidth * 0.5;
      var cy = cssHeight * 0.43;
      var minSide = Math.min(cssWidth, cssHeight);
      var radius = minSide * (0.19 + smooth.bass * 0.055 + smooth.pulse * 0.025);
      var rotY = t * (0.42 + smooth.mid * 0.7);
      var rotX = -0.38 + Math.sin(t * 0.4) * 0.16;
      var hue = scene.hue + smooth.bass * 65;

      ctx.shadowBlur = 10 + smooth.volume * 16;
      ctx.shadowColor = hsla(hue, 100, 60, 0.34);

      var latStep = scene.quality < 0.72 ? 2 : 1;
      for (var lat = -5; lat <= 5; lat += latStep) {
        var theta = (lat / 6) * Math.PI * 0.5;
        var ringRadius = Math.cos(theta) * radius;
        var y = Math.sin(theta) * radius * Math.cos(rotX);
        ctx.beginPath();
        ctx.ellipse(cx, cy + y, ringRadius, ringRadius * (0.25 + Math.abs(Math.sin(rotX)) * 0.16), rotY * 0.3, 0, Math.PI * 2);
        ctx.strokeStyle = hsla(hue + lat * 10, 95, 62, 0.11 + smooth.volume * 0.14);
        ctx.lineWidth = 1 + smooth.bass * 2;
        ctx.stroke();
      }

      var longitudes = scene.quality < 0.72 ? 7 : 10;
      for (var lon = 0; lon < longitudes; lon += 1) {
        drawLongitude(cx, cy, radius, rotX, rotY + lon * Math.PI / longitudes, hue + lon * 12);
      }

      drawSpherePoints(cx, cy, radius, rotX, rotY, hue);
      drawCore(cx, cy, radius * 0.23, t);
    }

    function drawLongitude(cx, cy, radius, rotX, rotY, hue) {
      var steps = scene.quality < 0.7 ? 64 : 88;
      ctx.beginPath();
      for (var i = 0; i <= steps; i += 1) {
        var theta = (i / steps) * Math.PI * 2;
        var wobble = 1 + sampleFrequency(i / steps) * 0.16 + smooth.pulse * 0.035;
        var x = Math.cos(theta) * Math.cos(rotY) * radius * wobble;
        var z = Math.cos(theta) * Math.sin(rotY) * radius * wobble;
        var y = Math.sin(theta) * radius * wobble;
        var rotatedY = y * Math.cos(rotX) - z * Math.sin(rotX);
        var rotatedZ = y * Math.sin(rotX) + z * Math.cos(rotX);
        var scale = 560 / (560 + rotatedZ);
        var px = cx + x * scale;
        var py = cy + rotatedY * scale;

        if (i === 0) {
          ctx.moveTo(px, py);
        } else {
          ctx.lineTo(px, py);
        }
      }
      ctx.strokeStyle = hsla(hue, 96, 62, 0.12 + smooth.volume * 0.16);
      ctx.lineWidth = 1.1 + smooth.high * 1.8;
      ctx.stroke();
    }

    function drawSpherePoints(cx, cy, radius, rotX, rotY, hue) {
      var step = scene.quality < 0.68 ? 6 : 4;
      for (var i = 0; i < scene.particles.length; i += step) {
        var p = scene.particles[i];
        var theta = p.angle + scene.phase * (0.45 + p.speed * 0.3);
        var phi = p.lane;
        var audioPush = 1 + sampleFrequency((i % 96) / 96) * 0.22 + smooth.pulse * 0.08;
        var x = Math.cos(theta) * Math.sin(phi) * radius * audioPush;
        var y = Math.sin(theta) * Math.sin(phi) * radius * audioPush;
        var z = Math.cos(phi) * radius * audioPush;
        var rx = x * Math.cos(rotY) - z * Math.sin(rotY);
        var rz = x * Math.sin(rotY) + z * Math.cos(rotY);
        var ry = y * Math.cos(rotX) - rz * Math.sin(rotX);
        rz = y * Math.sin(rotX) + rz * Math.cos(rotX);
        var scale = 560 / (560 + rz);
        var size = (1.2 + p.size * 2.5 + smooth.high * 4) * scale;
        drawDot(cx + rx * scale, cy + ry * scale, size, hsla(hue + p.hue, 98, 66, 0.22 + scale * 0.22));
      }
    }

    function drawAurora(t) {
      var baseHue = scene.hue;
      var layers = config.performance ? 3 : Math.max(4, Math.round(5 * scene.quality));
      var centerY = cssHeight * 0.42;

      drawParticleField(cssWidth * 0.5, centerY, t, 16, 0.35, false);

      for (var layer = 0; layer < layers; layer += 1) {
        var yBase = centerY + (layer - layers / 2) * cssHeight * 0.07;
        var amp = cssHeight * (0.045 + layer * 0.012) * (1 + smooth.bass * 1.2);
        var hue = baseHue + layer * 28 + smooth.high * 80;
        var gradient = ctx.createLinearGradient(0, yBase - amp * 2, cssWidth, yBase + amp * 2);
        gradient.addColorStop(0, hsla(hue, 96, 62, 0));
        gradient.addColorStop(0.28, hsla(hue, 96, 62, 0.34));
        gradient.addColorStop(0.62, hsla(hue + 70, 94, 62, 0.3));
        gradient.addColorStop(1, hsla(hue + 120, 94, 62, 0));

        ctx.beginPath();
        for (var x = 0; x <= cssWidth; x += Math.max(9, cssWidth / 140)) {
          var ratio = x / cssWidth;
          var freq = sampleFrequency(Math.pow(ratio, 1.25));
          var wave = Math.sin(ratio * Math.PI * (2.8 + layer * 0.22) + t * (0.7 + layer * 0.09));
          var ripple = Math.sin(ratio * Math.PI * 8 + t * (1.3 + smooth.mid));
          var y = yBase + wave * amp + ripple * amp * 0.22 + (freq - 0.22) * amp * 1.9;
          if (x === 0) {
            ctx.moveTo(x, y);
          } else {
            ctx.lineTo(x, y);
          }
        }
        ctx.strokeStyle = gradient;
        ctx.lineWidth = 3.2 + layer * 0.65 + smooth.volume * 5;
        ctx.shadowBlur = 4 + smooth.volume * 10;
        ctx.shadowColor = hsla(hue, 96, 62, 0.52);
        ctx.stroke();
      }

      drawCore(cssWidth * 0.5, centerY, Math.min(cssWidth, cssHeight) * (0.045 + smooth.bass * 0.035), t);
    }

    function drawSpectrum(t) {
      var columns = config.performance ? 42 : Math.round(68 * scene.quality);
      var gap = Math.max(2, cssWidth / columns * 0.16);
      var colW = cssWidth / columns - gap;
      var floor = cssHeight * 0.86;
      var maxH = cssHeight * 0.54;

      for (var i = 0; i < columns; i += 1) {
        var ratio = i / (columns - 1);
        var freq = sampleFrequency(Math.pow(ratio, 1.18));
        var eased = Math.pow(freq, 0.72);
        var h = Math.max(5, eased * maxH * (0.72 + config.sensitivity * 0.24) + smooth.bass * 28);
        var x = i * (colW + gap);
        var hue = scene.hue + ratio * 160 + smooth.bass * 50;
        var gradient = ctx.createLinearGradient(0, floor - h, 0, floor);
        gradient.addColorStop(0, hsla(hue + 40, 98, 72, 0.96));
        gradient.addColorStop(0.52, hsla(hue, 98, 60, 0.62));
        gradient.addColorStop(1, hsla(hue - 35, 94, 48, 0.14));

        ctx.fillStyle = gradient;
        ctx.shadowBlur = 0;
        roundRect(ctx, x, floor - h, Math.max(2, colW), h, Math.min(6, colW * 0.4));
        ctx.fill();

        ctx.globalAlpha = 0.22;
        roundRect(ctx, x, floor + 8, Math.max(2, colW), h * 0.28, Math.min(6, colW * 0.4));
        ctx.fill();
        ctx.globalAlpha = 1;
      }

      drawRadialBars(
        cssWidth * 0.5,
        cssHeight * 0.37,
        Math.min(cssWidth, cssHeight) * 0.08,
        Math.min(cssWidth, cssHeight) * 0.24,
        t,
        Math.round((config.performance ? 42 : 58) * scene.quality),
        0.48
      );
    }

    function drawParticleField(cx, cy, t, dt, yScale, vortex) {
      var minSide = Math.min(cssWidth, cssHeight);
      var depth = 980;
      var focal = 620;

      var step = scene.quality < 0.64 ? 2 : 1;
      for (var i = 0; i < scene.particles.length; i += step) {
        var p = scene.particles[i];
        p.z -= dt * (0.04 + p.speed * 0.035) * (1 + smooth.bass * 2.2);
        if (p.z < -depth * 0.5) {
          p.z += depth;
          p.angle = Math.random() * Math.PI * 2;
        }

        var scale = focal / (focal + p.z);
        var freq = sampleFrequency((i % 128) / 128);
        var radius = p.radius * minSide * (0.52 + smooth.bass * 0.12 + freq * 0.1);
        var angle = p.angle + scene.phase * (vortex ? 1.8 + p.speed : 0.35 + p.speed * 0.25);
        var wobble = Math.sin(t * (0.8 + p.speed) + p.wobble) * minSide * 0.018 * (1 + smooth.high);
        var x = cx + Math.cos(angle) * radius * scale + wobble;
        var y = cy + Math.sin(angle) * radius * yScale * scale + Math.cos(angle * 2 + t) * smooth.mid * 34;
        var size = (p.size * 1.18 + freq * 4.6 + smooth.pulse * 3.2) * scale;
        var alpha = clamp(0.08 + scale * 0.28 + freq * 0.34 + smooth.volume * 0.14, 0.06, 0.78);
        drawDot(x, y, size, hsla(scene.hue + p.hue + freq * 90, 96, 64, alpha));
      }
    }

    function drawRadialBars(cx, cy, inner, outer, t, count, yScale) {
      var rotation = t * (0.12 + smooth.mid * 0.16);
      ctx.shadowBlur = 0;

      for (var i = 0; i < count; i += 1) {
        var ratio = i / count;
        var freq = sampleFrequency(Math.pow(ratio, 1.12));
        var angle = ratio * Math.PI * 2 + rotation;
        var length = (outer - inner) * (0.24 + Math.pow(freq, 0.68)) + smooth.pulse * 30;
        var start = inner * (1 + smooth.bass * 0.05);
        var end = start + length;
        var x1 = cx + Math.cos(angle) * start;
        var y1 = cy + Math.sin(angle) * start * yScale;
        var x2 = cx + Math.cos(angle) * end;
        var y2 = cy + Math.sin(angle) * end * yScale;
        var hue = scene.hue + ratio * 210 + smooth.high * 90;

        ctx.beginPath();
        ctx.moveTo(x1, y1);
        ctx.lineTo(x2, y2);
        ctx.strokeStyle = hsla(hue, 98, 66, 0.24 + freq * 0.72);
        ctx.lineWidth = 1.6 + freq * 6.4 + smooth.pulse * 2.4;
        ctx.stroke();
      }
    }

    function drawCore(cx, cy, radius, t) {
      var hue = scene.hue + smooth.bass * 70;
      var pulseRadius = radius * (1 + smooth.volume * 1.2 + smooth.pulse * 0.9);
      var gradient = ctx.createRadialGradient(cx, cy, 0, cx, cy, pulseRadius * 3.2);
      gradient.addColorStop(0, hsla(hue + 20, 98, 82, 0.78));
      gradient.addColorStop(0.28, hsla(hue, 96, 62, 0.46));
      gradient.addColorStop(1, "rgba(0,0,0,0)");

      ctx.fillStyle = gradient;
      ctx.shadowBlur = 10 + smooth.pulse * 12;
      ctx.shadowColor = hsla(hue, 98, 64, 0.7);
      ctx.beginPath();
      ctx.arc(cx, cy, pulseRadius * 3.2, 0, Math.PI * 2);
      ctx.fill();

      for (var i = 0; i < 4; i += 1) {
        var r = pulseRadius * (1.2 + i * 0.58 + smooth.pulse * 0.45);
        ctx.beginPath();
        ctx.ellipse(cx, cy, r * (1 + i * 0.12), r * (0.36 + i * 0.04), t * (0.35 + i * 0.12), 0, Math.PI * 2);
        ctx.strokeStyle = hsla(hue + i * 34, 98, 68, 0.34 - i * 0.045);
        ctx.lineWidth = 1.8 + smooth.bass * 2.8;
        ctx.stroke();
      }
    }

    function spawnSparks() {
      var cx = cssWidth * 0.5;
      var cy = cssHeight * 0.43;
      var minSide = Math.min(cssWidth, cssHeight);

      scene.rings.push({
        x: cx,
        y: cy,
        radius: minSide * (0.08 + smooth.bass * 0.04),
        speed: minSide * (0.0016 + smooth.volume * 0.0028),
        life: 1,
        hue: scene.hue + 18 + smooth.high * 90
      });

      if (scene.rings.length > 10) {
        scene.rings.splice(0, scene.rings.length - 10);
      }

      if (config.performance) {
        return;
      }

      var count = Math.round((6 + smooth.bass * 16) * scene.quality);

      for (var i = 0; i < count; i += 1) {
        var angle = Math.random() * Math.PI * 2;
        var speed = 1.6 + Math.random() * 5.8;
        scene.sparks.push({
          x: cx,
          y: cy,
          vx: Math.cos(angle) * speed,
          vy: Math.sin(angle) * speed * 0.66,
          life: 1,
          size: 1.2 + Math.random() * 3.2,
          hue: scene.hue + Math.random() * 90
        });
      }

      if (scene.sparks.length > 96) {
        scene.sparks.splice(0, scene.sparks.length - 96);
      }
    }

    function drawBeatRings(dt) {
      for (var i = scene.rings.length - 1; i >= 0; i -= 1) {
        var ring = scene.rings[i];
        ring.radius += ring.speed * dt * 34;
        ring.life -= dt * 0.00115;

        if (ring.life <= 0) {
          scene.rings.splice(i, 1);
          continue;
        }

        ctx.beginPath();
        ctx.ellipse(ring.x, ring.y, ring.radius, ring.radius * 0.56, scene.phase * 0.32, 0, Math.PI * 2);
        ctx.strokeStyle = hsla(ring.hue, 100, 68, ring.life * 0.46);
        ctx.lineWidth = 2 + ring.life * 8;
        ctx.shadowBlur = 0;
        ctx.stroke();

        ctx.beginPath();
        ctx.ellipse(ring.x, ring.y, ring.radius * 0.72, ring.radius * 0.4, -scene.phase * 0.28, 0, Math.PI * 2);
        ctx.strokeStyle = hsla(ring.hue + 72, 100, 68, ring.life * 0.22);
        ctx.lineWidth = 1.5 + ring.life * 4;
        ctx.stroke();
      }
    }

    function drawSparks(dt) {
      for (var i = scene.sparks.length - 1; i >= 0; i -= 1) {
        var s = scene.sparks[i];
        s.x += s.vx * dt * 0.06;
        s.y += s.vy * dt * 0.06;
        s.vx *= 0.986;
        s.vy *= 0.986;
        s.life -= dt * 0.0016;

        if (s.life <= 0) {
          scene.sparks.splice(i, 1);
          continue;
        }

        drawDot(s.x, s.y, s.size * (0.5 + s.life), hsla(s.hue, 100, 66, s.life * 0.42));
      }
    }

    function drawDot(x, y, size, color) {
      if (size <= 0.2 || x < -40 || x > cssWidth + 40 || y < -40 || y > cssHeight + 40) {
        return;
      }

      ctx.fillStyle = color;
      ctx.shadowBlur = 0;
      if (size > 2.2) {
        ctx.globalAlpha *= 0.26;
        ctx.beginPath();
        ctx.arc(x, y, size * 2.9, 0, Math.PI * 2);
        ctx.fill();
        ctx.globalAlpha /= 0.26;
      }
      ctx.beginPath();
      ctx.arc(x, y, size, 0, Math.PI * 2);
      ctx.fill();
    }

    function createParticles() {
      var count = Math.round((config.performance ? 76 : 150) * config.density);
      var particles = [];

      for (var i = 0; i < count; i += 1) {
        particles.push({
          angle: Math.random() * Math.PI * 2,
          radius: Math.pow(Math.random(), 0.58) * 0.88 + 0.04,
          z: (Math.random() - 0.5) * 980,
          speed: 0.35 + Math.random() * 1.9,
          size: 0.85 + Math.random() * 2.7,
          hue: Math.random() * 180 - 50,
          lane: Math.acos(2 * Math.random() - 1),
          wobble: Math.random() * Math.PI * 2
        });
      }

      scene.particles = particles;
      scene.sparks = [];
      scene.rings = [];
    }

    function sampleFrequency(ratio) {
      if (!audio.frequency || audio.frequency.length === 0) {
        return smooth.volume * 0.15;
      }
      var index = clamp(Math.floor(ratio * (audio.frequency.length - 1)), 0, audio.frequency.length - 1);
      var value = audio.frequency[index] / 255;
      return clamp(Math.pow(value * config.sensitivity, 0.72), 0, 1);
    }

    function updatePanel(now, force) {
      if (!force && now - lastPanelAt < 140) {
        return;
      }
      lastPanelAt = now;

      var hasMedia = !!audio.media;
      var connected = !!audio.source;
      var waitingUnlock = audio.context && audio.context.state === "suspended";
      var sourceLabel = connected ? audio.sourceKind : hasMedia ? "attente" : "media ?";
      var status = "";

      if (!hasMedia) {
        status = "Lecteur introuvable. Lance une musique sur YouTube Music.";
      } else if (isPlaybackStopped()) {
        status = "Pause. Visualiseur en veille.";
      } else if (waitingUnlock) {
        status = "Clique dans la page pour deverrouiller l'analyse audio.";
      } else if (!connected) {
        status = audio.error || "Connexion audio en attente.";
      } else if (!metrics.hasSignal) {
        status = "Connecte, en attente de signal audible.";
      } else {
        status = "Signal OK. Analyse basses, volume et rythme.";
      }

      if (performance.now() < scene.reconnectNoticeUntil) {
        status = "Reconnexion audio...";
      }

      panel.classList.toggle("has-signal", !!metrics.hasSignal);
      panel.classList.toggle("no-media", !hasMedia);
      elements.source.textContent = sourceLabel;
      elements.status.textContent = status;
      elements.bass.style.width = Math.round(clamp(smooth.bass * 100, 0, 100)) + "%";
      elements.volume.style.width = Math.round(clamp(smooth.volume * 100, 0, 100)) + "%";
      elements.bpm.textContent = metrics.bpm ? String(metrics.bpm) : "--";
    }

    function destroy() {
      stopLoop();
      cleanup.forEach(function (dispose) {
        try {
          dispose();
        } catch (error) {
          // Keep destroying even if one listener is already gone.
        }
      });
      cleanup = [];
      disconnectCurrentSource();
      if (host) {
        host.remove();
      }
      if (window.__YTMV_APP__ === window.YTMVisualizer) {
        delete window.__YTMV_APP__;
      }
      if (window.YTMVisualizer && window.YTMVisualizer.destroy === destroy) {
        delete window.YTMVisualizer;
      }
    }

    function roundRect(context, x, y, width, height, radius) {
      var r = Math.min(radius, width / 2, height / 2);
      context.beginPath();
      context.moveTo(x + r, y);
      context.arcTo(x + width, y, x + width, y + height, r);
      context.arcTo(x + width, y + height, x, y + height, r);
      context.arcTo(x, y + height, x, y, r);
      context.arcTo(x, y, x + width, y, r);
      context.closePath();
    }

    function assign(target) {
      for (var i = 1; i < arguments.length; i += 1) {
        var source = arguments[i] || {};
        Object.keys(source).forEach(function (key) {
          target[key] = source[key];
        });
      }
      return target;
    }

    function clamp(value, min, max) {
      return Math.min(max, Math.max(min, value));
    }

    function hsla(h, s, l, a) {
      return "hsla(" + Math.round((h % 360 + 360) % 360) + "," + s + "%," + l + "%," + a + ")";
    }
  }
})();
