/* =========================================
   1. SYSTEM INITIALIZATION & CURSOR
   ========================================= */
const cursor = document.getElementById('cursor');
const introOverlay = document.getElementById('intro-overlay');
const startBtn = document.getElementById('start-btn');

// Mouse Movement for Custom Cursor
document.addEventListener('mousemove', (e) => {
    cursor.style.left = e.clientX + 'px';
    cursor.style.top = e.clientY + 'px';
});

// Hover effects on clickable elements
const clickables = document.querySelectorAll('a, button, .glass-panel, i');
clickables.forEach(el => {
    el.addEventListener('mouseenter', () => cursor.classList.add('active'));
    el.addEventListener('mouseleave', () => cursor.classList.remove('active'));
});

// Start Button Logic
startBtn.addEventListener('click', () => {
    introOverlay.style.opacity = '0';
    setTimeout(() => {
        introOverlay.style.visibility = 'hidden';
    }, 1000);
});

/* =========================================
   2. TYPEWRITER EFFECT
   ========================================= */
const words = [
    "PYTHON DEVELOPER",
    "SYSTEM ARCHITECT",
    "AUTOMATION SPECIALIST",
    "WEB DEVELOPER"
];
let i = 0, timer;

function typeWriter(text, n) {
    if (n < (text.length)) {
        document.getElementById("typewriter").innerHTML = text.substring(0, n+1) + '<span aria-hidden="true">|</span>';
        setTimeout(() => typeWriter(text, n + 1), 100);
    } else {
        setTimeout(() => eraseWriter(text, n), 2000);
    }
}

function eraseWriter(text, n) {
    if (n >= 0) {
        document.getElementById("typewriter").innerHTML = text.substring(0, n) + '<span aria-hidden="true">|</span>';
        setTimeout(() => eraseWriter(text, n - 1), 50);
    } else {
        i = (i + 1) % words.length;
        setTimeout(() => typeWriter(words[i], 0), 500);
    }
}

// Start typewriter after load
setTimeout(() => typeWriter(words[0], 0), 2000);

/* =========================================
   4. MATRIX "BLACK HOLE" PHYSICS ENGINE
   ========================================= */
const canvas = document.getElementById('matrix-canvas');
const ctx = canvas.getContext('2d');
let width, height;
let particles = [];

// Configuration
const particleCount = 150;
const mouse = { x: -1000, y: -1000 };

function resize() {
    width = canvas.width = window.innerWidth;
    height = canvas.height = window.innerHeight;
}
window.addEventListener('resize', resize);
resize();

window.addEventListener('mousemove', e => {
    mouse.x = e.clientX;
    mouse.y = e.clientY;
});

class Particle {
    constructor() {
        this.reset();
    }

    reset() {
        this.x = Math.random() * width;
        this.y = Math.random() * height;
        this.val = Math.random() > 0.5 ? '1' : '0';
        this.size = Math.random() * 10 + 8;
        this.vx = (Math.random() - 0.5) * 0.5;
        this.vy = (Math.random() - 0.5) * 0.5;
        this.baseColor = '#0f3';
    }

    update() {
        const dx = mouse.x - this.x;
        const dy = mouse.y - this.y;
        const distSq = dx*dx + dy*dy;
        const forceRadius = 25000;

        if (distSq < forceRadius) {
            const dist = Math.sqrt(distSq);
            const force = (forceRadius - distSq) / forceRadius;
            const angle = Math.atan2(dy, dx);

            this.vx += Math.cos(angle) * force * 0.8;
            this.vy += Math.sin(angle) * force * 0.8;
        } else {
            if (this.vx > 1) this.vx *= 0.95;
            if (this.vx < -1) this.vx *= 0.95;
            if (this.vy > 1) this.vy *= 0.95;
            if (this.vy < -1) this.vy *= 0.95;
        }

        this.x += this.vx;
        this.y += this.vy;

        if (this.x < 0) this.x = width;
        if (this.x > width) this.x = 0;
        if (this.y < 0) this.y = height;
        if (this.y > height) this.y = 0;

        if (Math.random() < 0.02) this.val = Math.random() > 0.5 ? '1' : '0';
    }

    draw() {
        ctx.font = `${this.size}px monospace`;

        const dx = mouse.x - this.x;
        const dy = mouse.y - this.y;
        const dist = Math.sqrt(dx*dx + dy*dy);

        if (dist < 150) {
            ctx.fillStyle = `rgba(255, 255, 255, ${1 - dist/150})`;
        } else {
            ctx.fillStyle = Math.random() > 0.9 ? '#fff' : 'rgba(0, 247, 255, 0.3)';
        }

        ctx.fillText(this.val, this.x, this.y);
    }
}

for(let i=0; i<particleCount; i++) {
    particles.push(new Particle());
}

function animateMatrix() {
    ctx.fillStyle = 'rgba(5, 5, 5, 0.2)';
    ctx.fillRect(0, 0, width, height);

    particles.forEach(p => {
        p.update();
        p.draw();
    });
    requestAnimationFrame(animateMatrix);
}
animateMatrix();

/* =========================================
   5. CAROUSEL 3D EFFECT
   ========================================= */
document.addEventListener('DOMContentLoaded', function() {
    const carousels = document.querySelectorAll('.carousel');

    carousels.forEach(carousel => {
        const items = carousel.querySelectorAll('.carousel-item');
        const numItems = items.length;

        if (numItems < 3) {
            console.error('Le carousel doit avoir au moins 3 éléments.');
            return;
        }

        const angle = 360 / numItems;
        const radius = 150; // Rayon du cercle sur lequel les éléments sont placés

        items.forEach((item, index) => {
            const currentAngle = angle * index;
            const x = radius * Math.sin(currentAngle * Math.PI / 180);
            const z = radius * Math.cos(currentAngle * Math.PI / 180);

            item.style.transform = `rotateY(${currentAngle}deg) translateZ(${z}px) translateX(${x}px)`;
        });
    });
});

// Recalculer la géométrie du carousel lors d'un resize
window.addEventListener('resize', function() {
    const carousels = document.querySelectorAll('.carousel');

    carousels.forEach(carousel => {
        const items = carousel.querySelectorAll('.carousel-item');
        const numItems = items.length;

        if (numItems < 3) {
            console.error('Le carousel doit avoir au moins 3 éléments.');
            return;
        }

        const angle = 360 / numItems;
        const radius = 150; // Rayon du cercle sur lequel les éléments sont placés

        items.forEach((item, index) => {
            const currentAngle = angle * index;
            const x = radius * Math.sin(currentAngle * Math.PI / 180);
            const z = radius * Math.cos(currentAngle * Math.PI / 180);

            item.style.transform = `rotateY(${currentAngle}deg) translateZ(${z}px) translateX(${x}px)`;
        });
    });
});

// Arrêter l'animation de rotation lors du hover
const carouselItems = document.querySelectorAll('.carousel-item');

carouselItems.forEach(item => {
    item.addEventListener('mouseenter', function() {
        this.parentElement.style.animationPlayState = 'paused';
    });

    item.addEventListener('mouseleave', function() {
        this.parentElement.style.animationPlayState = 'running';
    });
});
