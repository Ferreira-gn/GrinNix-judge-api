// Este arquivo é o "test harness" — nunca muda, só o USER_CODE muda
const userCode   = process.env.USER_CODE;
const fnName     = process.env.FUNCTION_NAME || "solution";
const testCases  = JSON.parse(process.env.TEST_CASES || "[]");

// Executa o código do usuário no escopo global
try {
  eval(userCode);
} catch (err) {
  const results = testCases.map((_, i) => ({
    index: i,
    passed: false,
    error: `Syntax/runtime error: ${err.message}`,
  }));
  console.log(JSON.stringify(results));
  process.exit(0);
}

// Pega a função pelo nome
const fn = global[fnName] ?? eval(fnName);

if (typeof fn !== "function") {
  console.log(JSON.stringify([{ index: 0, passed: false, error: `Function "${fnName}" not found` }]));
  process.exit(0);
}

// Roda cada test case
const results = testCases.map((tc, index) => {
  try {
    const args   = Array.isArray(tc.input) ? tc.input : [tc.input];
    const output = fn(...args);

    const passed =
      JSON.stringify(output) === JSON.stringify(tc.expected);

    return {
      index,
      passed,
      output,
      expected: tc.expected,
    };
  } catch (err) {
    return {
      index,
      passed: false,
      error: err.message,
      expected: tc.expected,
    };
  }
});

console.log(JSON.stringify(results));