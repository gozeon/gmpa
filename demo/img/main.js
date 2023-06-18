import bmp from './20x20-1KB.bmp'
import png from './20x20-1KB.png'
import jpg from './20x20-1KB.jpg'
import svg from './test.svg'

const insertImg = url => {
    let image = new Image
    image.src = url
    document.body.appendChild(image)
}

insertImg(bmp)
insertImg(png)
insertImg(jpg)
insertImg(svg)

