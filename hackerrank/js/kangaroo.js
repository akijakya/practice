function kangaroo(x1, v1, x2, v2) {
    x1 += v1;
    x2 += v2;
    if (x1 === x2) {
        return 'YES';
    } else if (x1 > x2 || v1 < v2) {
        return 'NO';
    } else {
        return kangaroo(x1, v1, x2, v2);
    }
}

console.log(kangaroo(0, 3, 4, 2));
console.log(kangaroo(0, 5, 5, 1));