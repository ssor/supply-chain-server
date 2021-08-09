import { North, East, South, West } from './direction.js'
import { PositionLine } from './PositionLine.js'

export default class Truck {
	direction = East;//0 1 2 3 -> N  E  N W
	sprite;
	// destX = [];
	positionLine
	currentGoal = null;
	goalReached = false;
	number
	text
	speed

	constructor(number, direction) {
		// this.text = text;
		this.positionLine = new PositionLine();
		this.number = number;
		this.direction = direction;
	}

	addDest(p) {
		console.log(`add dest ->`);
		// console.table(p)
		if (Array.isArray(p)) {
			for (let index = 0; index < p.length; index++) {
				const element = p[index];
				this.positionLine.AddPosition(p)
			}
			// this.destX.push(...p);
		} else {
			this.positionLine.AddPosition(p)
			// let c = this.destX.length;
			// if (c > 0) {
			// 	let last = this.destX[c - 1]
			// 	if (this.moreOrLessEqual(last.x, p.x) && this.moreOrLessEqual(last.y, p.y)) {
			// 		// duplicate no need to add
			// 	} else {
			// 		this.destX.push(p);
			// 		console.info("add dest OK, ", this.destX.length)
			// 	}
			// } else {
			// 	this.destX.push(p);
			// }
		}
	}
	moreOrLessEqual(f, s) {
		if (Math.abs(f - s) <= 3) {
			return true
		}
		return false
	}

	setPosition(p) {
		this.sprite.x = p.x;
		this.sprite.y = p.y;
	}
	setSpeed(s) {
		this.speed = s;
	}
	move(direction) {
		switch (direction) {
			case North:
				break;
			case East:
				console.log(`move -> ${direction}`)
				if (this.direction !== East) {
					this.turnAround(East);
				}
				this.sprite.setVelocityX(this.speed);
				break;
			case South:

				break;
			case West:
				if (this.direction !== West) {
					this.turnAround(West);
				}
				this.sprite.setVelocityX(-this.speed);
				break;

			default:
				console.error(`no direction ${direction}`)
				break;
		}
	}

	calculateDirection(x, dest) {
		if (x <= dest.x) {
			console.log(`${x} -> ${dest.x} = East`)
			return East;
		} else {
			console.log(`${x} -> ${dest.x} = West`)
			return West;
		}
	}

	nextGoal() {
		// let next = this.destX.shift();
		let next = this.positionLine.Shift();
		if (next.maybe == false) {
			// console.log('no next goal to move')
			this.sprite.setVelocityX(0);//stop
			return
		}
		console.info("next: ", next);
		this.currentGoal = next.position;
		console.info("new goal set,", this.currentGoal);
		// console.info("left dest count: ", this.destX.length);
		this.goalReached = false;

		const direction = this.calculateDirection(this.sprite.x, this.currentGoal)
		this.turnAround(direction);
		switch (direction) {
			case North:
				break;
			case East:
				console.log(`move -> ${direction}`)
				this.sprite.setVelocityX(this.speed);
				break;
			case South:

				break;
			case West:
				this.sprite.setVelocityX(-this.speed);
				break;

			default:
				console.error(`no direction ${direction}`)
				break;
		}
	}

	turnAround(d) {
		switch (d) {
			case East:
				this.direction = East;
				this.sprite.setRotation(Math.PI / 2);
				// this.sprite.setAngle(90);
				console.log("turn to East")
				break;
			case West:
				this.direction = West;
				// this.sprite.setAngle(-90);
				this.sprite.setRotation(-Math.PI / 2);
				console.log("turn to West")
				break;
		}
	}

	checkGoal() {
		if (this.currentGoal == null) {
			this.goalReached = true;
		}
		if (this.goalReached === true) {
			this.nextGoal()
		} else {
			const x = this.sprite.x;
			if (Math.abs(x - this.currentGoal.x) <= 3) {
				//reach goal
				// nextGoal();
				this.goalReached = true;
			} else {
				// not yet reached goal
			}
		}
	}

	updateNumberTextPosition() {
		this.text.x = this.sprite.x;
		this.text.y = this.sprite.y;
	}
	checkBoundary(maxWidth) {
		this.checkGoal();

		this.updateNumberTextPosition();

		const width = this.sprite.displayWidth;
		const max = maxWidth - width;
		switch (this.direction) {
			case East:
				if (this.sprite.x > max) {
					this.sprite.setVelocityX(0);
				} else {
					// console.log(`truck now at ${this.sprite.x} -> ${maxWidth}`);
				}
				break;
			case West:
				if (this.sprite.x <= width) {
					this.sprite.setVelocityX(0);
				}
				break;
		}
	}

}