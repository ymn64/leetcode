const join = require('./join');

const tests = [
  {
    name: 'test 1',
    arr1: [
      { id: 1, x: 1 },
      { id: 2, x: 9 },
    ],
    arr2: [{ id: 3, x: 5 }],
    want: [
      { id: 1, x: 1 },
      { id: 2, x: 9 },
      { id: 3, x: 5 },
    ],
  },
  {
    name: 'test 2',
    arr1: [
      { id: 1, x: 2, y: 3 },
      { id: 2, x: 3, y: 6 },
    ],
    arr2: [
      { id: 2, x: 10, y: 20 },
      { id: 3, x: 0, y: 0 },
    ],
    want: [
      { id: 1, x: 2, y: 3 },
      { id: 2, x: 10, y: 20 },
      { id: 3, x: 0, y: 0 },
    ],
  },
  {
    name: 'test 3',
    arr1: [{ id: 1, b: { b: 94 }, v: [4, 3], y: 48 }],
    arr2: [{ id: 1, b: { c: 84 }, v: [1, 3] }],
    want: [{ id: 1, b: { c: 84 }, v: [1, 3], y: 48 }],
  },
];

for (const t of tests) {
  test(t.name, () => {
    const got = join(t.arr1, t.arr2);
    expect(got).toEqual(t.want);
  });
}

test(tests[0].name, () => {
  const got = join(tests[0].arr1, tests[0].arr2);
  expect(got).toEqual(tests[0].want);
});
