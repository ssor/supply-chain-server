import Phaser from './phaser.js'
// import Truck from './truck.js'

import Game from './Game.js'


var gameScene = new Game();

var game = new Phaser.Game({
	type: Phaser.AUTO,
	width: 1280,
	height: 720,
	// canvas: document.getElementById("gameCanvas"),
	physics: {
		default: 'arcade',
		arcade: {
			gravity: {
				y: 0
			},
			debug: true,
		}
	},
	scene: gameScene,
	fps: {
		target: 10
	}
});


export {game, gameScene}
    // var game = new Phaser.Game(config);
    // truck.addDest(keyPoints);
    // truck.addDest(keyPoints);