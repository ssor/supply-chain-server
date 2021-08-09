import Phaser from './phaser.js'
import Truck from './Truck.js'
import Producer from './producer.js'
import { North, East, South, West } from './direction.js'

export default class Game extends Phaser.Scene {
	text1
	cursors
	truck
	producer

	constructor() {
		super("game");
		this.truck = new Truck(21, East);

	}

	preload() {
		this.load.image('bg', '/public/assets/img/sc-bg2.png');
		this.load.image('truck', '/public/assets/img/truck.png');
		this.load.image('producer-1', '/public/assets/img/producer-1.png');
		this.load.image('producer-2', '/public/assets/img/producer-2.png');
		this.load.image('dispatcher1-1', '/public/assets/img/dispatcher1-1.png');
		this.load.image('dispatcher1-2', '/public/assets/img/dispatcher1-2.png');
		this.load.image('dispatcher2-1', '/public/assets/img/dispatcher2-1.png');
		this.load.image('dispatcher2-2', '/public/assets/img/dispatcher2-2.png');
		this.load.image('detailer-1', '/public/assets/img/detailer1.png');
		this.load.image('detailer-2', '/public/assets/img/detailer2.png');
		this.load.image('gg1', '/public/assets/img/green-group-1.png');
		this.load.image('gg2', '/public/assets/img/green-group-2.png');
		this.load.image('gg3', '/public/assets/img/green-group-3.png');
	}

	create() {

		this.cursors = this.input.keyboard.createCursorKeys();

		this.add.image(640, 360, 'bg');
		this.add.image(100, 103, 'producer-1').setScale(0.6, 0.6);
		this.add.image(400, 103, 'dispatcher1-1').setScale(0.6, 0.6);
		this.add.image(860, 103, 'dispatcher2-1').setScale(0.6, 0.6);
		this.add.image(1200, 103, 'detailer-1').setScale(0.6, 0.6);

		for (let index = 0; index < 4; index++) {
			this.add.image(80 + 145 * index, 300, 'gg1').setScale(0.6, 0.6);
		}
		for (let index = 0; index < 4; index++) {
			this.add.image(770 + 145 * index, 300, 'gg2').setScale(0.6, 0.6);
		}
		for (let index = 0; index < 4; index++) {
			this.add.image(80 + 145 * index, 530, 'gg3').setScale(0.6, 0.6);
		}
		for (let index = 0; index < 4; index++) {
			this.add.image(770 + 145 * index, 530, 'gg1').setScale(0.6, 0.6);
		}

		this.text1 = this.add.text(10, 10, '', { fill: '#ff0000', font: "bold 30px" });
		
		{
			let x = 100;
			let y = 50;
			let styleInventory = { fill: '#ffffff', font: "bold 18px Arial" }
			let styleNumber = { fill: '#ffffff', font: "12px Arial" }
			let numberText = this.make.text({
				x: x, y: y - 25, origin: 0.5, style: styleNumber
			});
			let inventoryText = this.make.text({
				x: x, y: y, origin: 0.5, style: styleInventory
			});
			this.producer = new Producer("01", x, y, numberText, inventoryText, 100);
		}
		{
			
			var truckSprite = this.physics.add.sprite(0, 180, 'truck').setScale(0.3).setRotation(Math.PI / 2).setInteractive();
			let x = 200;
			let y = 50;
			let styleNumber = { fill: '#ffffff', font: "bold 16px Arial" }
			let numberText = this.make.text({
				x: x, y: y - 25, origin: 0.5, style: styleNumber
			});
			numberText.setText("21");
			// truckSprite.on('pointerdown', function (pointer) {
			// 	console.log("truck clicked, now  at %d", truckSprite.x);
			// });
			this.truck.sprite = truckSprite;
			this.truck.text = numberText;
		}

	}



	update() {
		var pointer = this.input.activePointer;

		this.text1.setText([
			'x: ' + pointer.worldX,
			'y: ' + pointer.worldY,
			'isDown: ' + pointer.isDown
		]);

		// console.log(`x: ${pointer.worldX} y: ${pointer.worldY}`)

		if (this.cursors.left.isDown) {
			this.truck.move(West);
		}
		else if (this.cursors.right.isDown) {
			this.truck.move(East);
		}

		if (this.cursors.up.isDown) {
			// player.setVelocityY(-300);
		}
		else if (this.cursors.down.isDown) {
			// player.setVelocityY(300);
		}
		// console.log(this.scale.width);
		this.truck.checkBoundary(this.scale.width);
	}
}
