/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-27
 * @fileoverview This program prints the Fibonacci sequence up to a user-entered limit.
 */

// variables
let limit: number = 0;
let firstNumber: number = 0;
let secondNumber: number = 1;
let nextNumber: number = 0;

// get limit from user
limit = parseInt(prompt("Enter the limit for Fibonacci series: ") || "0");

console.log("");
console.log(firstNumber);
console.log(secondNumber);

// while loop to generate Fibonacci numbers
nextNumber = firstNumber + secondNumber;

while (nextNumber <= limit) {
  console.log(nextNumber);

  firstNumber = secondNumber;
  secondNumber = nextNumber;
  nextNumber = firstNumber + secondNumber;
}

console.log("\nDone.");
