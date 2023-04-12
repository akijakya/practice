function breakingRecords(scores) {
    // let score = scores.split(' ').map(x => parseInt(x));
    let min = scores[0];
    let max = scores[0];
    let result = [0, 0]
    
    for (let i = 0; i < scores.length; i++) {
        if (scores[i] > max) {
            max = scores[i];
            result[0] += 1;
        } else if (scores[i] < min) {
            min = scores[i];
            result[1] += 1;
        }
    }

    return result
}

console.log(breakingRecords([10, 5, 20, 20, 4, 5, 2, 25, 1]));