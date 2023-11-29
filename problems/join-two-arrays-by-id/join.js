module.exports = function join(arr1, arr2) {
  const map = new Map(arr1.map((item) => [item.id, item]));

  for (const a of arr2) {
    if (map.has(a.id)) {
      Object.assign(map.get(a.id), a);
    } else {
      map.set(a.id, a);
    }
  }

  return Array.from(map.values()).sort((a, b) => a.id - b.id);
};
