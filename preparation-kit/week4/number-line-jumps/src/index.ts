function kangaroo(x1: number, v1: number, x2: number, v2: number): string {
	if (v1 === v2) {
		return "NO";
	}

	const jumps = (x2 - x1) / (v1 - v2);

	console.log(x2 - x1, v1 - v2, jumps);

	if (Number.isInteger(jumps) && jumps >= 0) {
		return "YES";
	}

	return "NO";
}

function main() {
	console.log(kangaroo(0, 2, 5, 3));
	console.log(kangaroo(0, 3, 4, 2));
	console.log(kangaroo(5, 1, 4, 2));
	console.log(kangaroo(11, 2, 2, 4));
}

main();

// 0 - 3 - 6 - 9 - 12
// 4 - 6 - 8 - 10 - 12
// ----
// 5 - 6
// 4 - 6
