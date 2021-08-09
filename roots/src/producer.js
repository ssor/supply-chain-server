

export default class Producer {
    numberText
    inventoryText
    x
    y
    maxInventory
    number
    constructor(number, x, y, numberText, inventoryText, maxInventory) {
        this.x = x;
        this.y = y;
        this.number = number;
        this.maxInventory = maxInventory;
        this.numberText = numberText;
        this.numberText.setText("生产商-"+this.number);
        this.inventoryText = inventoryText;
        this.inventoryText.setText("0 / " + maxInventory)
    }

    updateProductText(count) {
        this.inventoryText.setText(count + " / " + this.maxInventory)
    }
}