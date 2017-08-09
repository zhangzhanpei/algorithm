//计算陆地的边长
//一个由0和1组成的二维数组, 1表示陆地. 0与1相接的边需要计算, 矩形的边界也要计算
var islandPerimeter = function(grid) {
    let row = grid.length; //1
    let col = grid[0].length; //2
    let perimeter = 0;
    for (let i = 0; i < row; i++) {
        for (let j = 0; j < col; j++) {
            if (grid[i][j] === 0) {
                continue;
            }
            //上
            if (i + 1 === row || grid[i + 1][j] === 0) {
                perimeter += 1;
            }
            //右
            if (j + 1 === col || grid[i][j + 1] === 0) {
                perimeter += 1;
            }
            //下
            if (i === 0 || grid[i - 1][j] === 0) {
                perimeter += 1;
            }
            //左
            if (j === 0 || grid[i][j - 1] === 0) {
                perimeter += 1;
            }
        }
    }
    return perimeter;
};

console.log(islandPerimeter([[1, 0]])); //output 4
