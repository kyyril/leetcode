function possibleStringCount(word: string): number {
  let count = 1; // original string is always a valid possibility

  let i = 0;
  while (i < word.length) {
    let j = i;
    while (j < word.length && word[j] === word[i]) {
      j++;
    }
    const groupLen = j - i;
    if (groupLen > 1) {
      count += groupLen - 1;
    }
    i = j;
  }

  return count;
}
