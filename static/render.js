
function render(currentLocation, allObjects) {
    let canvas = document.getElementById("map");
    let context = canvas.getContext("2d");
    context.clearRect(0, 0, canvas.width, canvas.height);
    drawBorders(context, canvas.width, canvas.height, 3);
    for(let i in allObjects) {
        let obj = allObjects[i];
        if (obj["objectType"] === "Player") {
            drawPlayer(context, obj);
        }
        if (obj["objectType"] === "Bullet") {
            drawBullet(context, obj);
        }
    }
    console.log(currentLocation, allObjects)
}

function drawBorders(context, canvasWidth, canvasHeight, borderWidth) {
    context.fillStyle = "black";
    context.fillRect(0, 0, canvasWidth, borderWidth);
    context.fillRect(0, 0, borderWidth, canvasHeight);
    context.fillRect(0, canvasHeight - borderWidth, canvasWidth, canvasHeight);
    context.fillRect(canvasWidth - borderWidth, 0, canvasWidth, canvasHeight);
}

function drawPlayer(context, player) {
    context.fillStyle = "green";
    context.beginPath();
    context.arc(player["point"]["x"], player["point"]["y"], 25, 0, 2 * Math.PI);
    context.fill();
}

function drawBullet(context, bullet) {
    context.fillStyle = "yellow";
    context.fillRect(bullet["point"]["x"] - 5, bullet["point"]["y"] - 5, 10, 10);
}