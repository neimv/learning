// es no hice nada, es apra manejo de imagenes
const sharp = require("sharp");

sharp("original.png")
    .resize(80)
    .toFile("resized.png")