$(document).ready(function () {
    if(window["WebSocket"]) {
        let host = $("script[src*=game]").attr("data-host");
        let conn = new WebSocket("ws://" + host + "/ws");

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
        let canvas = document.getElementById("map")

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
            let data = JSON.parse(event.data);
            currentLocation = data["myLocation"];
            render(currentLocation, data["render"]);
        };

        setInterval(function () {
            conn.send(JSON.stringify({"movement": movement, "mouseLocation": mouseLocation, "shoot": shoot}));
            shoot = false;
        }, 1000 / 30)

    } else {
        alert("Cant connect by websocket")
    }
});