let canvas = document.getElementById('myCanvas');
let imageDataUrl = canvas.toDataURL();
fetch('https://target.io/upload', {
  method: 'POST',
  body: JSON.stringify({ imageDataUrl }),
  headers: { 'Content-Type': 'application/json' },
});
