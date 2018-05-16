function Vector(angle, length) {
    this.Angle = angle;
    this.Length = length;
}

function NewVector(x, y) {
    let length = Math.sqrt(x*x + y*y);
    if (length == 0) {
        return new Vector(0, 0);
    }
    return new Vector(Math.atan2(y, x), length);
}

Vector.prototype.GetX = function() {
    return Math.cos(this.Angle) * this.Length;
};

Vector.prototype.GetY = function () {
    return Math.sin(this.Angle) * this.Length;
};

Vector.prototype.Add = function (vector) {
   return NewVector(this.GetX() + vector.GetX(), this.GetY() + vector.GetY());
};

Vector.prototype.Normallize = function () {
    if (this.Length == 0)
        return new Vector(0, 0);
    return new Vector(this.Angle, 1);
};

Vector.prototype.MultiplyByScalar = function (scalar) {
    return new Vector(this.Angle, this.Length * scalar)
}
