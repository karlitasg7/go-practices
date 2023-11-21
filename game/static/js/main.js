
let imgArray = [
  "static/img/hand_rock.png",
  "static/img/hand_paper.png",
  "static/img/hand_scissors.png"
];

function choice(x) {
  fetch(`/play?c=${x}`)
    .then(response => response.json())
    .then(data => {

      if (x === 0) {
        document.getElementById("player_choice").innerHTML = "Player select ROCK"
      } else if (x === 1) {
        document.getElementById("player_choice").innerHTML = "Player select PAPER"
      } else {
        document.getElementById("player_choice").innerHTML = "Player select SCISSORS"
      }

      document.getElementById("player_score").innerHTML = data.player_score
      document.getElementById("computer_score").innerHTML = data.computer_score
      document.getElementById("computer_choice").innerHTML = `Computer select ${data.computer_choice}`
      document.getElementById("round_result").innerHTML = data.round_result
      document.getElementById("round_message").innerHTML = data.message

      document.getElementById("computer_img").src = imgArray[data.computer_choice_int]

    })
}