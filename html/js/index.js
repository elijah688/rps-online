
let playerScore = 0;
let cpuScore = 0;
const rps = async (choice) => {
    const {
        cpu_move,
        player_move,
        winner
    } = await playRound(choice)

    
    document.getElementById("winner-you").innerHTML = ""
    document.getElementById("winner-cpu").innerHTML = ""

    document.getElementById("chose-you").innerHTML = player_move
    document.getElementById("chose-cpu").innerHTML = cpu_move
    
    if (player_move === cpu_move) {
        document.getElementById("winner-you").innerHTML = winner
        document.getElementById("winner-cpu").innerHTML = winner
    } else if (winner === "Player Wins!") {
        playerScore++;
        document.getElementById("winner-you").innerHTML = winner
        document.getElementById("score-you").innerHTML = playerScore
    }
    else {
        cpuScore++;
        console.log(cpuScore)
        document.getElementById("winner-cpu").innerHTML = winner
        document.getElementById("score-cpu").innerHTML = cpuScore

    }
}

const playRound = async (choice) => {
    const response = await fetch(`/rps?rps=${choice}`, {
        method: 'GET',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json'
        },
    });
    const json = await response.json()
    return json;
}

// <td id="chose-you">@chose-you</td>
// <td id="winner-you">@winner-you</td>
// </tr>
// <tr>
// <th scope="row">CPU</th>
// <td id="chose-cpu">@chose-cpu</td>
// <td i