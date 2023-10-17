# Canvas GPU Side Channel

A tool for GPU side channel analysis using an HTML canvas rendered in a browser. The objective is to detect differences in canvas element rendering from different machines due to their different GPU driver. By rendering the same canvas data and rasterizing the image, then diffing the target image against a control image, the GPU and driver of the target image can be identified.

(work in progress)
