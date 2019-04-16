常规做法就是单循环里用n不断减去m，时间复杂度`O(n/m)`。可以采用类似矩阵快速幂的优化方法，
事先预处理出所有的`m*(2^n)<=n`,这样复杂度就能降低到log级

time complexity: `O(log(n/m))`

space complexity: `O(log(n/m))`