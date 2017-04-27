/*
 *插板法分配红包金额
 *100块分成30个红包, 先将100块每份1块分成100份, 排成一列, 然后随机位置插入29个插板, 将这列红包隔成随机的30份
 */
function randomMoney(money, n)
{
    let index = [];
    //生成 n - 1 个插板, 将 money 分成 n 份
    for (let i = 0; i < n - 1; i++) {
        let random = Math.floor(Math.random() * money);
        while (index.indexOf(random) >= 0) {
            random = Math.floor(Math.random() * money);
        }
        index.push(random);
        //插板位置排序
        index.sort(function(a, b){
            if (a < b) {
                return -1;
            } else if (a > b) {
                return 1;
            } else {
                return 0;
            }
        });
    }
    //根据插板位置算红包金额
    let moneyArr = [];
    moneyArr.push(index[0]); //第一个红包金额就是第一个插板的位置[第一个插板减去0位置的插板]
    for (let j = 1; j < n - 1; j++) {
        moneyArr.push(index[j] - index[j - 1]); //后一个插板的位置减去前一个插板的位置就是两个插板之间的金额
    }
    moneyArr.push(money - index[n - 2]); //最后一个红包的金额
    console.log(moneyArr);
}

randomMoney(100, 30);
