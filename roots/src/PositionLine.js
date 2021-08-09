import { Position, MaybePosition } from './Position.js'

class PositionLine {
    NextIndex
    SN
    Positions = []
    constructor() {
        this.SN = 0;
        this.Positions = [];
        this.NextIndex = 0;
    }

    Shift() {
        if (this.SN <= 0) {
            return new MaybePosition(false)
        }
        if (this.NextIndex >= this.SN) {
            return new MaybePosition(false)
        }
        let p = this.Positions[this.NextIndex]
        this.NextIndex++
        return new MaybePosition(true, p);
    }

    AddPosition(p) {
        if (this.SN <= 0) {
            this.Positions.push(p)
            this.SN++
            return
        }
        let last = this.Positions[this.SN - 1];
        if (last.sn >= p.sn) {
            return
        }
        this.Positions.push(p)
        this.SN++
    }

}
export { PositionLine }