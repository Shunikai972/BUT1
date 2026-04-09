# bing_first_image.py
from __future__ import annotations

import asyncio
import json
import re
from pathlib import Path
from typing import Optional
from urllib.parse import quote_plus, urlsplit
import mimetypes

import httpx
from playwright.async_api import async_playwright, Browser, BrowserContext, Page


class BingFirstImageDownloader:
    def __init__(
        self,
        headless: bool = True,
        timeout_ms: int = 15000,
        downloads_dir: str = "downloads",
        locale: str = "fr-FR",
    ):
        self.headless = headless
        self.timeout_ms = timeout_ms
        self.downloads_dir = Path(downloads_dir)
        self.downloads_dir.mkdir(parents=True, exist_ok=True)
        self.locale = locale

        self._playwright = None
        self._browser: Optional[Browser] = None
        self._context: Optional[BrowserContext] = None
        self._page: Optional[Page] = None

    async def start(self) -> None:
        if self._browser:
            return

        self._playwright = await async_playwright().start()
        self._browser = await self._playwright.chromium.launch(headless=self.headless)
        self._context = await self._browser.new_context(
            locale=self.locale,
            viewport={"width": 1400, "height": 1000},
            user_agent=(
                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
                "AppleWebKit/537.36 (KHTML, like Gecko) "
                "Chrome/124.0.0.0 Safari/537.36"
            ),
        )
        self._page = await self._context.new_page()
        self._page.set_default_timeout(self.timeout_ms)

    async def close(self) -> None:
        if self._context:
            await self._context.close()
        if self._browser:
            await self._browser.close()
        if self._playwright:
            await self._playwright.stop()

        self._playwright = None
        self._browser = None
        self._context = None
        self._page = None

    async def _accept_cookies_if_present(self) -> None:
        if not self._page:
            return

        possible_selectors = [
            'button:has-text("Accepter")',
            'button:has-text("Accept")',
            'input[value="Accept"]',
            '#bnp_btn_accept',
        ]
        for selector in possible_selectors:
            try:
                btn = self._page.locator(selector).first
                if await btn.is_visible(timeout=1200):
                    await btn.click()
                    return
            except Exception:
                pass

    async def _extract_first_image_url(self, query: str) -> str:
        if not self._page:
            raise RuntimeError("Browser not started")

        search_url = f"https://www.bing.com/images/search?q={quote_plus(query)}&form=HDRSC3"
        await self._page.goto(search_url, wait_until="domcontentloaded")
        await self._accept_cookies_if_present()

        # Attendre que les résultats d'images apparaissent
        card = self._page.locator("a.iusc").first
        await card.wait_for(state="visible")

        # Bing stocke souvent l'URL image dans l'attribut 'm' (JSON)
        m_attr = await card.get_attribute("m")
        if m_attr:
            try:
                data = json.loads(m_attr)
                for key in ("murl", "imgurl", "turl"):
                    value = data.get(key)
                    if value and value.startswith("http"):
                        return value
            except Exception:
                pass

        # Fallback : récupérer l'image visible dans la carte
        img = card.locator("img").first
        for attr in ("src", "data-src", "data-bm", "data-thumburl"):
            val = await img.get_attribute(attr)
            if val and val.startswith("http"):
                return val

        raise RuntimeError("Impossible d'extraire l'URL de la première image Bing.")

    @staticmethod
    def _safe_filename(name: str) -> str:
        name = re.sub(r"[\\/:*?\"<>|]+", "_", name)
        name = re.sub(r"\s+", " ", name).strip()
        return name[:140] or "image"

    @staticmethod
    def _guess_extension(url: str, content_type: Optional[str]) -> str:
        path = urlsplit(url).path.lower()

        for ext in (".jpg", ".jpeg", ".png", ".webp", ".gif", ".bmp"):
            if path.endswith(ext):
                return ext

        if content_type:
            ext = mimetypes.guess_extension(content_type.split(";")[0].strip())
            if ext:
                return ".jpg" if ext == ".jpe" else ext

        return ".jpg"

    async def download_first_bing_image(
        self,
        query: str,
        output_path: Optional[str] = None,
    ) -> dict:
        await self.start()

        image_url = await self._extract_first_image_url(query)

        async with httpx.AsyncClient(
            follow_redirects=True,
            timeout=20.0,
            headers={
                "User-Agent": (
                    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
                    "AppleWebKit/537.36 (KHTML, like Gecko) "
                    "Chrome/124.0.0.0 Safari/537.36"
                ),
                "Referer": "https://www.bing.com/",
            },
        ) as client:
            response = await client.get(image_url)
            response.raise_for_status()

            ext = self._guess_extension(image_url, response.headers.get("content-type"))

            if output_path:
                file_path = Path(output_path)
                if not file_path.suffix:
                    file_path = file_path.with_suffix(ext)
            else:
                file_path = self.downloads_dir / f"{self._safe_filename(query)}{ext}"

            file_path.parent.mkdir(parents=True, exist_ok=True)
            file_path.write_bytes(response.content)

        return {
            "query": query,
            "image_url": image_url,
            "saved_to": str(file_path.resolve()),
        }


# API simple pour un autre programme
_downloader: Optional[BingFirstImageDownloader] = None

async def get_downloader() -> BingFirstImageDownloader:
    global _downloader
    if _downloader is None:
        _downloader = BingFirstImageDownloader()
        await _downloader.start()
    return _downloader

async def download_first_bing_image(query: str, output_path: Optional[str] = None) -> dict:
    downloader = await get_downloader()
    return await downloader.download_first_bing_image(query, output_path)


# Mode CLI
async def _main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("query", help="Nom à rechercher sur Bing Images")
    parser.add_argument("-o", "--output", help="Chemin de sortie du fichier")
    args = parser.parse_args()

    downloader = BingFirstImageDownloader()
    try:
        result = await downloader.download_first_bing_image(args.query, args.output)
        print(json.dumps(result, ensure_ascii=False, indent=2))
    finally:
        await downloader.close()


if __name__ == "__main__":
    asyncio.run(_main())