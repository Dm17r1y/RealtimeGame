$(document).ready(function () {
    let attrs = $("script[src*=game]");
    let host = attrs.attr("data-host");
    let delay = attrs.attr("data-delay");
    let playerSpeed = attrs.attr("data-player-speed");
    let conn = new WebSocket("ws://" + host + "/ws");
    let currentState = {render: [], stateIndex: -1, tick: 0};
    let commands = new Queue();
    let tick = null;

    let movement = {
        up: false,
        down: false,
        left: false,
        right: false,
    };

    let mouseLocation = {
        x: 0,
        y: 0,
    };

    let shoot = false;
    let canvas = document.getElementById("map");

    document.addEventListener('keydown', function (event) {
        switch (event.key.toLowerCase()) {
            case "a":
                movement.left = true;
                break;
            case "w":
                movement.up = true;
                break;
            case "d":
                movement.right = true;
                break;
            case "s":
                movement.down = true;
                break;
        }
    });

    document.addEventListener('keyup', function (event) {
        switch (event.key.toLowerCase()) {
            case "a":
                movement.left = false;
                break;
            case "w":
                movement.up = false;
                break;
            case "d":
                movement.right = false;
                break;
            case "s":
                movement.down = false;
                break;
        }
    });

    canvas.addEventListener('mousedown', function (event) {
        shoot = true
    });

    canvas.addEventListener('mousemove', function (event) {
        mouseLocation = {"x": event.clientX, "y": event.clientY};
    });

    conn.onmessage = function (event) {
        let state = JSON.parse(event.data);
        while (commands.size() > 0 && commands.getData(0).tick < state.tick) {
            commands.dequeue()
        }

        for(let i = 0; state.tick + i < tick && i < commands.size(); i++) {

            state = ApplyCommand(state, commands.getData(i), playerSpeed);
        }
        tick = state.tick;
        currentState = state;
    };

    setInterval(function () {
        if (tick == null)
            return
        let command = {"movement": movement, "mouseLocation": mouseLocation, "shoot": shoot};
        conn.send(JSON.stringify(command));
        shoot = false;
        commands.enqueue({"command": command, "tick": tick});
        currentState = ApplyCommand(currentState, command, playerSpeed);
        render(currentState.render);
    }, delay)
});

function ApplyCommand(state, command, playerSpeed) {
    if (state.stateIndex === -1)
        return;
    let stateCopy = $.extend(true, {}, state);
    let player = stateCopy.render[state.stateIndex];
    let up = new Vector(-Math.PI / 2, 1);
    let down = new Vector(Math.PI / 2, 1);
    let left = new Vector(Math.PI, 1);
    let rigth = new Vector(0, 1);

    let movement = new Vector(0, 0);

    if (command.movement.left) {
        movement = movement.Add(left)
    }
    if (command.movement.right) {
        movement = movement.Add(rigth)
    }
    if (command.movement.up) {
        movement = movement.Add(up)
    }
    if (command.movement.down) {
        movement = movement.Add(down)
    }

    movement = movement.Normallize().MultiplyByScalar(playerSpeed);
    player.position.x += movement.GetX();
    player.position.y += movement.GetY();

    direction = NewVector(command.mouseLocation.x - player.position.x, command.mouseLocation.y - player.position.y);
    player.direction = direction.Angle;
    return stateCopy;
}