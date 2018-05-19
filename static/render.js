
function render(allObjects) {
    let canvas = document.getElementById("map");
    let context = canvas.getContext("2d");
    context.clearRect(0, 0, canvas.width, canvas.height);
    drawBorders(context, canvas.width, canvas.height, 3);
    for(let i in allObjects) {
        let obj = allObjects[i];
        if (obj.objectType === "Player") {
            drawPlayer(context, obj);
        }
        if (obj.objectType === "Bullet") {
            drawBullet(context, obj);
        }
        if (obj.objectType === "FastBullet") {
            drawFastBullet(context, obj)
        }
    }
}

function drawFastBullet(ctx, bullet) {
    ctx.beginPath();
    ctx.moveTo(bullet.position.x, bullet.position.y);
    endPosition = NewVector(bullet.position.x, bullet.position.y)
        .Add(new Vector(bullet.direction, 1)
        .MultiplyByScalar(1000));
    ctx.lineTo(endPosition.GetX(), endPosition.GetY());
    ctx.strokeStyle = "blue";
    ctx.stroke()
}

function drawBorders(ctx, canvasWidth, canvasHeight, borderWidth) {
    ctx.fillStyle = "black";
    ctx.fillRect(0, 0, canvasWidth, borderWidth);
    ctx.fillRect(0, 0, borderWidth, canvasHeight);
    ctx.fillRect(0, canvasHeight - borderWidth, canvasWidth, canvasHeight);
    ctx.fillRect(canvasWidth - borderWidth, 0, canvasWidth, canvasHeight);
}

function drawPlayer(ctx, player) {
    drawPlayerBody(ctx, player.position.x, player.position.y);
    drawPlayerGun(ctx, player.position.x, player.position.y, 25, player.direction)
}

function drawPlayerBody(ctx, x, y) {
    ctx.beginPath();
    ctx.arc(x, y, 25, 0, 2 * Math.PI);
    ctx.fillStyle = "green";
    ctx.fill();
}

function drawPlayerGun(ctx, x, y, length, rotationAngle) {
    // ctx.save();
    // ctx.beginPath();
    // ctx.translate(x, y);
    // ctx.rotate(rotationAngle - Math.PI / 2);
    // ctx.rect(-width / 2, 0, width, height);
    // ctx.fillStyle = "red";
    // ctx.fill();
    // ctx.restore();
    ctx.beginPath();
    ctx.moveTo(x, y);
    ctx.lineTo(x + Math.cos(rotationAngle) * length, y + Math.sin(rotationAngle) * length);
    ctx.strokeStyle = "red";
    ctx.stroke()
}

function drawBullet(ctx, bullet) {
    ctx.fillStyle = "yellow";
    ctx.fillRect(bullet.position.x - 5, bullet.position.y - 5, 10, 10);
}