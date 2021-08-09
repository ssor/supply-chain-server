
class Position {
	x
	y
	sn
	constructor(sn, x, y) {
		this.sn = sn;
		this.x = x;
		this.y = y;
	}
   	
}

class MaybePosition {
	maybe
	position
	constructor(maybe, p) {
		this.maybe = maybe;
		this.position = p;
	}
}

export { Position, MaybePosition }