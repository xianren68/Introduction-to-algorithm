听b站大神左程云算法课做的笔记

https://www.bilibili.com/video/BV13g41157hK/?p=7&spm_id_from=333.1007.top_right_bar_window_history.content.click
老师讲过的题目，如果在力扣上有，我也会带上链接，方便大家练习
### 复杂度
### 位运算
#### 异或
> 异或运算的性质
> 1. 相同的数异或为0，一个数和0异或为它本身
> 2. 异或运算满足交换率和结合率

**注意：go中没有按位与的符号，可以通过^n来实现**
**取某个数n的最右边的1**`rightone := n&(^n+1)`
###### 拓展
```
1. 一个数组，只有一个数为奇数个，其他均为偶数个，求这个数
```
leetcode:https://leetcode.cn/problems/single-number/
> https://github.com/xianren68/Introduction-to-algorithm/tree/main/bitwise/01_异或.go
```
2. 一个数组中，有两个数位奇数个，求这两个数
```


### 排序
#### 1.选择排序
> 从0开始进行n次循环，每次循环从当前值到数组末位，选择出一个最大或最小的值与当前位置进行交换
> 时间复杂度O(n^2)，空间复杂度o(1)

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/basic_sort/1.选择排序.go
#### 2.冒泡排序
> 与选择排序相似，每次循环找出最小或最大的值，将其移动到数组末位，不过选择排序会记录最小或最大的值
> 一次循环只进行一次交换，而冒泡排序则对相邻元素比较，交换
> 时间复杂度o(n^2),空间复杂度o(1)

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/basic_sort/2.冒泡排序.go
#### 3.插入排序
> 将要排序的值插入到已排序的值中，从第二值个开始遍历数组，将当前值前面的数作为已排序的序列，将当前值插入其中，
> 从而让有序序列的长度再次加1，然后再继续向下遍历,相比较于前面两个排序，插入排序的性能因为数据而有所不同，
> 当数据较为有序时，它的遍历次数会小很多，而选择与冒泡不会发生变化
> 时间复杂度o(n^2),空间复杂度o(1)

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/basic_sort/3.插入排序.go
#### 4.归并排序
> 通过递归的方式，将给定的数据分为左右两个序列，分别让它们有序，然后将其合并，形成一个新的有序序列
> 时间复杂度,根据master公式为O(N*logN)，空间复杂度为O(N)

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/basic_sort/4.归并排序.go
###### 拓展
```
1.求一个数组的小和 数组小和的定义如下：
例如，数组s = [1, 3, 5, 2, 4, 6]，在s[0]的左边小于或等于s[0]的数的和为0；在s[1]的左边小于或等于s[1]的数的和为1；在s[2]的左边小于或等于s[2]的数的和为1+3=4；在s[3]的左边小于或等于s[3]的数的和为1；
在s[4]的左边小于或等于s[4]的数的和为1+3+2=6；在s[5]的左边小于或等于s[5]的数的和为1+3+5+2+4=15。所以s的小和为0+1+4+1+6+15=27
给定一个数组s，实现函数返回s的小和
2.求数组逆序对
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对
```
#### 5.快速排序
> 1.荷兰国旗问题：给定一个数组，并给定一个值num，将比num大的数放到数组前段，将=num的数放到数组中段，大于
> num的数放到数组后段
> 解法：
> 1. 给定两个指针l,r,用来指定<区的后一位和>区的前一位
> 2. 通过索引i遍历数组,如果小于num,将当前位置与l位置交换，扩充<区，l++
> 3. 如果小于num,将当前位置与l位置交换，扩充>区，r--,不知道交换回值的大小，所以i应保持不变
> 4. 当i=r时，所有的数已经排列完毕
###### 快速排序2.0：
>1. 将数组的最后一位作为指定数，然后通过递归将数组转化为三块，再将指定数放到等于区
>2. 通过递归将数组各个部分都执行上述操作，最后得到有序的数组
###### 快速排序3.0：
>因为2.0版本的快排很受数据的影响，时间复杂度为O(n^2),3.0基于2.0将指定数num从最后一位，换成随机取，然后再与最后
>一位进行交换，因为随机，所以不会受数据状况影响，时间复杂度为O(n * logN),空间复杂度为O(logN)

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/basic_sort/5.快速排序.go
#### 6.堆排序
##### 堆（heap）
> 一个存储了一棵完全二叉树的数组
> 如果每个子树的根节点都比其子节点大，则称为大根堆，根节点是整个树中的最大值，反之则为小根堆
##### 性质
> 1. 当前值在数组中的索引为i，则其左子节点的索引为（i*2)+1,右子节点为（i*2)+2
> 2. 当前值在数组中的索引为i,其父节点的索引为（i-1)/2
##### 操作（以大根堆为例）
> 1. 将数据入堆，heapsize +1 ,数据加入到二叉数的最后一个节点，然后判断其与父节点的大小，若大，则交换位置，
> 重复这个过程，直到没有父节点或父节点比它大，
> 2. 将数据出堆，弹出数组首位最大值，二叉树的最后一个节点提到头节点的位置，然后将heapsize - 1,
> 头节点判断其左右子树中最大的一个是否比自己大，若是，则交换位置，然后重复此操作，直到没有子节点或字节点都比小
##### 堆排序
> 1. 将指定数组遍历，让每个值都经历一次入堆操作，得到一个大根堆
> 2. 再次遍历大根堆，每次将头节点交换到堆的最后，然后再通过出堆的操作让剩余的数据依然形成大根堆，头节点为最大值
> 3. 重复二操作，直至堆中没有数据，此时数组成为升序序列
> 时间复杂度O(N*logN),空间复杂度O(1)

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/basic_sort/6.堆排序.go
###### 拓展
```
已知一个几乎有序的数组，几乎有序是指一个数距离它排好序的距离不超过k,并且k相对数组长度较小，请选择合适的排序算法
进行排序
```
#### 7. 桶排序
> 前面的排序方式都是通过比较来进行排序的，我们也可以不进行比较来排序
##### 1.计数排序
> 首先，知道数组值的范围，定义出一个此范围的辅助数组，遍历原数组，原数组中的每个数都对应辅助数组上的一位
> 所以，可以统计出每个数出现的次数，然后根据辅助数组记录的计录的次数将数据填入源数组接口完成排序
> 时间复杂度O(N)，空间复杂度O(N)
##### 2.基数排序（桶排序）
> 根据数据的位数来进行排序，
> 1. 判断数据位数（十百千这种），根据数据的进制准备一个长度等于进制数的数组
> 2. 从最低位开始,将它们从小到大装到准备好的数组中，得到一个最低位升序的序列，将其放回原数组
> 3. 然后上升一位，再对原数组重复2操作，得到当前位得有序序列，将其放回原数组，一直重复上述操作，直到最高位也有序

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/basic_sort/7.基数排序.go
**注意：不通过比较的排序对数据的要求很高，如果数组中数的范围很大，就不适用，而比较排序可以适用于任何数组**
```
marst公式
在计算涉及递归的算法的时候，计算复杂度就会变得有些麻烦。Master公式就是用来进行剖析递归行为和递归行为时间复杂度的估算的

Master公式：T(N) = a*T(N/b) + O(N^d)

公式解释：n表示问题的规模，a表示递归的次数也就是生成的子问题数，N/b表示子问题的规模。O(N^d)表示除了递归操作以外其余操作的复杂度

结论（证明省略）：
①当d<logb a时，时间复杂度为O(N^(logb a))
②当d=logb a时，时间复杂度为O((N^d)*logN)
③当d>logb a时，时间复杂度为O(N^d)

注意：子问题规模必须等分，不管你是分成几部分
```
#### 排序算法稳定性
> 排序之后相等的值与排序下之前相对位置保持不变

|排序|时间复杂度|空间复杂度|稳定性|
|---|---|---|---|
|选择排序|O(N^2)|O(1)|不稳定|
|冒泡排序|O(N^2)|O(1)|相等值不换位置可以稳定|
|插入排序|O(N^2)|O(1)|相等值不换位置可以稳定|
|归并排序|O(NlogN)|O(N)|合并数组时相等时，先将左边的加入数组，稳定|
|快速排序|O(NlogN)|O(logN)|不稳定|
|堆排序|O(NlogN)|O(1)|不稳定|

#### 排序总结
> 1. 不基于比较的排序，对样本数据的要求高，不易改写
> 2. 基于比较的排序，只要规定好如何比较就可以直接复用
> 3. 基于比较的排序，时间复杂度最小为O(N*logN)
> 4. 绝对速度选快排，节省空间用堆排，稳定选归并


### .查找
#### 1.二分法
> 通常用来在有序集合中查找某个值，将数组为三，当前值，左边的，右边的
> 如果当前值是要找的值，则直接返回，否则判断左右哪个符合条件，再将其一分为3继续重复上述过程
> 最后总能找到要查找的值，时间复杂度O(logN)
> 其实只要能确定要查找的值必然在数组分隔的两端
> 就可以用二分法

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/search/01_二分法.go
###### 拓展
```
一个数组，求局部最小值(左右都比当前值大)
```

###  表
##### 有序表
> 其中所有元素以递增或递减方式有序排列
##### 哈希表
> 顺序存储的结构类型需要一个一个地按顺序访问元素，当这个总量很大且我们所要访问的元素比较靠后时，性能就会很低
> hash通过键值对映射的方式来加快访问记录，它是无序的
#### 链表
> 链表有一系列节点组成，节点主要包括当前节点值和它下一个节点的位置，由此链接成的链式结构
> 相比于数组，它的插入操作更快，但获取操作较慢
##### 重要技巧
> 笔试，以时间复杂度为要求
> 面试， 需要考虑时间复杂度
###### 常用方法
>1. 通过额外空间记录
>2. 快慢指针
###### 题目
```
1. 判断一个链表是否为回文链表
```
leetcode:https://leetcode.cn/problems/palindrome-linked-list/
> https://github.com/xianren68/Introduction-to-algorithm/tree/main/link_list/1.回文链表.go
```
2. 给定一个值，将链表分为小于给定值，等于给定值，大于给定值的三部分
```
leetcode:https://leetcode.cn/problems/partition-list-lcci/
> https://github.com/xianren68/Introduction-to-algorithm/tree/main/link_list/2.链表分区.go
```

3. 定义一种特殊的链表节点
{
    value,
    rand,
    next
}
rand是一个指针，可能指向某个节点或为空，给定一个由此节点构成的无环单链表，请将其复制，并返回头节点
```
leetcode:https://leetcode.cn/problems/copy-list-with-random-pointer/comments/
> https://github.com/xianren68/Introduction-to-algorithm/tree/main/link_list/3.复制链表.go
```
4.相交链表，给定两个链表，它们可能有环也可能没环，如果它们相交，返回它们相交的第一个节点，如果不相交，返回null
```
###### 题解
**切记千万不要脑部链表结构，最好画一下，看它们能不能用单链表实现**

**前置知识**
> 1. 求链表是否有环？
> 单链表有环只能有一种类似6的结构，环永远在最后
> 解法：
> 1. 通过hash表记录每个节点，在遍历的时候查看hash表中是否有此值，若有，则有环，并且第一个重复的值为环的起点
> 2. **快慢指针，如果有环，快慢指针总会相遇，如果相遇则有环，这时让慢指针停留到原地，快指针从头结点再开始跑，它们再次
> 相遇的地方就是入环的节点** （记住这个结论）

leetcode:https://leetcode.cn/problems/c32eOV/
<<<<<<< HEAD
>https://github.com/xianren68/Introduction-to-algorithm/tree/main/link_list/3.有环链表.go
=======
>https://github.com/xianren68/Introduction-to-algorithm/tree/main/link_list/4.有环链表.go
>>>>>>> 9635bd9a43d6c339b17cd4b874ddfcc01af6dd54

**链表相交**
> 1. 判断两个链表是否有环
> 2. 如果都没有环，则两个都遍历到末位，判断是否相等，若不相等，则不相交，若相等，记录它们的长度，让较长的先走它们的长度差距步，然后两个同时开始走，相遇的地方即为相交的地方
> 因为如果相交，它们的尾节点一定是同一个
> 3. 如果一个有环，一个没有环，则它们必不可能相交
> 若它们都有环，则先记录两个链表的入环节点，若它们相同，则可以将其看作都没有环，将入环节点看作尾节点，用2即可解决
> 如不相同，则可以让一个入环节点在环上转一圈，如果碰到另一个则说明两个链表相交，并且他们俩两个的入环节点都可以看作第一个相交的节点
> 如果没碰到，则二者没有相交

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/link_list/5.相交链表.go
### 二叉树 
#### 遍历
##### 递归 
###### 1.理解递归序
```
    func Traversal(root *ListNode){
        if (root == nil){
            return
        }
        // 1
        Traversal(root.Left)
        // 2
        Traversal(root.Right)
        // 3
    }
```
> 如上代码可以看到三个位置，函数在执行时会来到一个节点三次，在1，2，3的位置对当前的节点进行操作就是三种递归操作
> 1. 先序 根-左-右
> 2. 中序 左-根-右
> 3. 后序 左-右-根

leetcode:
先序：https://leetcode.cn/problems/binary-tree-preorder-traversal/
中序：https://leetcode.cn/problems/binary-tree-inorder-traversal/
后序：https://leetcode.cn/problems/binary-tree-postorder-traversal/

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/1.递归遍历.go
##### 非递归
######  1.先序遍历
> 准备一个栈，将头节点入栈，然后出栈时将它的子节点入栈，入栈时先入右子节点，然后继续重复上述出栈与入栈的操作，直到栈为空，弹出值的顺序为先序
###### 2. 后序遍历
> 准备两个栈，第一个栈操作与先序遍历一致，不过先入左节点，然后弹出的值压入另一个栈中，等第一个栈空了后开始将第二个栈的值依次弹出，直到栈为空，弹出值的顺序为后序
###### 3. 中序遍历
> 1. 准备一个栈，顺着左子节点直到将每个左节点入栈,直到整棵树的最左边节点也入栈
> 2. 从栈中弹出一个值，判断其是否有右节点，若有则进入右节点，重复1操作
> 3. 若没有右节点，重复1，2操作
> 弹出值的顺序为中序

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/2.非递归.go

##### 层序遍历
> 借助一个队列，将节点入队，出队时它的左右子节点入队，重复此过程，直到队列为空，返回的值即为层序

leetcode:https://leetcode.cn/problems/binary-tree-level-order-traversal/

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/3.层序遍历.go

###### 拓展
```
求二叉树最大宽度
```
leetcode:https://leetcode.cn/problems/maximum-width-of-binary-tree/

##### 常见的几种二叉树
###### 1.搜索二叉树
> 左子树的值必须小于当前值，右子树的值必须大于当前值
###### 解法：
> 中序遍历二叉搜索树，可以得到一个递增的序列
> 1. 用一个全局变量存储每一个节点的值，通过中序遍历来获取每个值，判断其是否大于上一个值
> 2. 如果它大于则继续判断下一个节点，并将值存到全局变量中
> 3. 如果它小于等于，它就不是搜索二叉树

leetcode:https://leetcode.cn/problems/validate-binary-search-tree/submissions/

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/4.二叉搜索树.go

###### 2.完全二叉树
> 按照顺序由左到右依次变满的树

###### 解法：
> 1. 通过层序遍历每个节点
> 2. 判断这个节点是否存在右节点而不存在左节点，若是，则不是完全二叉树
> 3. 如果一个节点只有左孩子而没有右孩子，或者它为叶子节点，则它后面的所有的节点都必须为叶子节点

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/5.完全二叉树.go

###### 3.满二叉树
> 一个二叉树，如果每一个层的结点数都达到最大值，则这个二叉树就是满二叉树。也就是说，如果一个二叉树的深度为K，且结点总数是(2^k) -1 ，则它就是满二叉树。

###### 解法：
> 1. 通过层序遍历得到最大深度与总节点个数，直接判断
> 2. 通过递归判断左右节点是否都为完全二叉树，如果是，则返回true

>https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/6.满二叉树.go

###### 4.平衡二叉树
> 某二叉树中任意节点的左右子树的深度相差不超过1则为平衡二叉树

###### 解法：
> 当前值通过递归拿到左右子树的深度进行判断

leetcode:https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/submissions/

>https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/7.平衡二叉树.go

###### 二叉树递归套路
> 判断当前值是否需要从左右子树获得数据，另左右子树返回的数据类型一致，再通过当前值判断或汇总
> 1. 搜索二叉树
> 从左子树获得其最小值并获得左子树是否为搜索二叉树，从右子树或的其最大值并获得左子树是否为搜索二叉树,两边的数据不一致，
> 如果要用套路递归，则可以让每个子树返回三个值，（最大值，最小值，是否是搜索二叉树）,统一数据后即可递归
> 满二叉树（返回子树是否为满的）以及平衡二叉树（返回子树的深度，子树是否为平衡二叉树）的判断用了这个套路
> **这个套路虽然不是适用于任何情况，但适用性很广，很多和树有关的问题都可以试试这钟做法**

###### 题目
```
1. 求两个节点的最近公共祖先节点
```
###### 解法：
**1.hash表**
>1. 通过hash表记录每个节点得父节点，通过父节点再找父节点，就可以形成一个祖先节点链
>2. 将一个节点得祖先节点链抽离出来放到另一个hash表中，
>3. 另一个节点通过hash表在祖先链上回溯，判断祖先节点是否在新抽离出的hash表中，若有，则直接返回
**2.递归**
> 两个节点在一棵树中，只有三种情况，要么在某棵子树的两侧，要么一个为另一个的父节点
> 递归时，碰到这两个值直接返回，（若它是另一个值的祖先节点，则返回的就是它，若它不是另一个值祖先节点，则说明它们必然
> 存在于某棵子树的两侧，它们得公共祖先节点在上面，直接返回让上面判断）
> 然后通过父节点判断其所对应的另一侧子树有无另一个值，若没有
> 则继续往上返回，若有，则说明它们就在这棵子树的两侧，直接返回当前子树的根节点即可

leetcode:https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

> https://github.com/xianren68/Introduction-to-algorithm/tree/main/binary_tree/7.最近祖先.go

###### 拓展
```
有一个如下节点组成的二叉树
type TreeNode struct {
    Val int  
    Left *TreeNode
    Right *TreeNode
    Parent *TreeNode
}
parent指向节点的父节点，根节点的parent为null,给定一个二叉树中的值，求它的后继节点
后继节点：中序遍历当前节点后面的节点为后继节点
```

```
将一纸条对折一次，会出现一个凹下去的折痕，对折两次纸条上的折痕为凹凹凸，三次为凹凹凸凹凹凸凸
给定一个值k，求其折痕分布
```

```
二叉树序列化与反序列化
以_作为一个值的结束，以#作为空值，然后将其链接成一个字符串即可
反序列化，直接根据_将其分隔成数组
```

### 图
> 图是由顶点集合以及顶点间的关系集合组成的一种数据结构。
> 即节点以及节点间线段连接成的数据结构
> 通常分为有向图（节点与节点之间的线段是有向的）
![alt:"有向图"](img/map/有向图.png)
> 以及无向图
![atl:"无向图"](img/map/无向图.png)
**出度与入度**
> 从某个节点出去的线的总数称为节点的出度,如上方有向图
> A的出度为2，进入某个节点的线的总数称为入度，A的入度为1
> 无向图的线可以看作是一条入的线和一条出的线组成的，所以它的出度和入度总是相等的
**邻接**
>可以从当前节点通过线走到的节点，即为当前节点的邻接节点，如上面有向图中A的邻接节点为B和D，因为是有向图，它无法到C
**权重**
> 两个节点之间的边的长度即为权重，没有权重的图称为无权图，上方的两个图都为无权图


#### 图的存储方式

> 1. 邻接表法
> 将每个节点存到容器中，每个节点又包含其邻接的所有节点，若有权重，则可以把每条边的权重也加上
![alt](img/map/邻接表.png)
> data为相邻的节点，weight为它们之间边的权重
> 2. 邻接矩阵
> 用所有节点组成一个矩阵，它们交叉的点的值即为它们之间的距离
> 没有边可以直接到达，距离即为无穷大
![alt](img/map/邻接矩阵.png)

#### 图的模板
> 图的方式可能千奇百怪，要解决它们的问题就得研究这每一种结构
> 非常的麻烦，我们可以定义一种存储图的模板，碰到图有关的问题
> 直接将它转化为我们熟悉的模板来解决它

[图模板](utils/mapTemplate.go)

```
一个n*3的矩阵，
[[0,1,2][2，4，5][3，4，6]]
每一层代表一个边，
第一位代表来的节点，第二位代表去的节点，第三位代表权重
将其转化为我们的模板
```
[转化矩阵](map_test/1.转化矩阵.go)
> 我封装了一个输出模板的方法，输入`{{1, 4, 5}, {2, 3, 6}, {7, 5, 3}}`，这样一个矩阵
> 组成的模板输出为
```
{值：3,入度：1,出度:0,邻接节点:
{值：7,入度：0,出度:1,邻接节点:（节点值：5,权重：3)}
{值：5,入度：1,出度:0,邻接节点:
{值：1,入度：0,出度:1,邻接节点:（节点值：4,权重：5)}
{值：4,入度：1,出度:0,邻接节点:
{值：2,入度：0,出度:1,邻接节点:（节点值：3,权重：6)}
```
> 有了图的模板，我们直接在模板上练习各种操作，如果碰到其他类型的图，将其转化为模板类型，即可继续操作
#### 图的遍历
##### 1.广度优先
> 1. 准备一个队列
> 2. 将源节点入队
> 3. 出队一个节点，将其所有还未进过队列的邻接节点入队（可以通过另一个hash表来判断邻接节点是否入过队）
> 4. 重复3，直到队列为空

[广度优先](map_test/2.广度优先.go)
##### 2.深度优先
> 1. 准备一个栈
> 2. 将源节点入栈
> 3. 弹出一个节点，并将其未进过栈的节点压入栈
> 4. 重复3,直到栈为空

[深度优先](map_test/3.深度优先.go)

##### 拓扑排序
> 一个有向无环图，从一个节点开始将所有顶点组成一个线性的序列
> 序列中每个顶点只能出现一次
> 任何一对由边连起来的顶点，线段出发的顶点总在前面
###### 应用
> 很多打包工具在打包依赖时就是拓扑排序，从一个入口文件开始，根据文件相互依赖的关系来将它们打包在一起
###### 写法
> 1. 准备一个队列存储入度为0的节点
> 2. 找到入度为0的节点，它就是入口
> 3. 将入度为0的节点的邻接节点的入度-1(注意：不要在图结构上直接改，可以用一个（节点->剩余入度）组成的hash表记录下来，然后在这个表中改)，然后将里面入度为0的节点加入入度为0的队列
> 4. 弹出一个入度为0的节点，重复上述过程，直到队列为空 
> 这个过程有点类似于宽度优先，但它们并不相同，这里的值只有入度为0才能加入队列
##### 生成树
>一个连通图（如果图中任意两点都是连通的，那么图被称作连通图。如果此图是有向图，则称为强连通图（注意：需要双向都有路径））的生成树是该连通图的一个极小连同子图,它含有图中全部顶点,和构成一棵树的(n-1)条边.如果在一棵生成树上添加任何一条边,必定构成一个环,因为这条边使得它依附的那两个顶点之间有了第二条路径.一棵有n个顶点的生成树(连通无回路图)有且仅有(n-1)条边,但是,有(n-1)条边的图不一定都是生成树.如果一个图有n个顶点和小于(n-1)条边,则是非连通图;如果它有多于(n-1)条边,则一定有回路.
###### 最小生成树
**适用于无向图**
>  对于一个带权(假定每条边上的权值均为大于零的实数)连通无向图G中的不同生成树,各树的边上的权值之和可能不同；图中所有生成树中具有边上的权值之和最小的树称为该图的最小生成树.
>按照生成树的定义,n个顶点的连通图的生成树有n个顶点和(n-1)条边.因此构造最小生成树的准则有三条:
> 1. 必须只使用该图中的边来构造最小生成树;
> 2. 必须使用且仅使用(n-1)条边来连接图中的n个顶点;
> 3. 不能使用产生回路的边.
###### 1.Kruskal
>   克鲁斯卡尔算法的基本思想是以边为主导地位，始终选择当前可用（所选的边不能构成回路）的最小权植边。所以Kruskal算法的第一步是给所有的边按照从小到大的顺序排序。这一步可以直接使用库函数qsort或者sort。接下来从小到大依次考察每一条边（u，v）
> 1. 获取到所有边中最小的边，判断边的两边是否是同一个集合，若是则重新挑选边
> 2. 若不是，则将边两边的集合合并，再继续找下一个最小边
> 3. 重复上述过程，直到没有边
> 因为要用到并查集，所以我没写全，等学到并查集再补全
###### 2.prime 
> 1. 准备一个小根堆
> 2. 从一个顶点出发，将它的所有出去的边加入到堆中，然后从堆中取出一个最小的边，将边的另一头加入到生成树中
> 3. 边另一头的节点也有许多边出去，将它们都加入到堆中，然后再从中取出最小的边，重复上面过程
> 4. 知道最小生成树的节点与原树的一样（遍历n-1次，每次拿到一个边，n为节点个数）
[最小生成树](/map_test/5.最小生成树.go)

参考：https://blog.csdn.net/qq_41181772/article/details/89357461?ops_request_misc=&request_id=&biz_id=102&utm_term=%E6%9C%80%E5%B0%8F%E7%94%9F%E6%88%90%E6%A0%91&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-0-89357461.142^v73^insert_down3,201^v4^add_ask,239^v2^insert_chatgpt&spm=1018.2226.3001.4187


##### Dijstra
**权重不能为负数**
> 求出一个有权图中一个顶点到其他所有顶点的最短路径

推荐 https://blog.csdn.net/qq_44431690/article/details/108175827

https://blog.csdn.net/lbperfect123/article/details/84281300?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-84281300-blog-108175827.pc_relevant_multi_platform_whitelistv4&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-84281300-blog-108175827.pc_relevant_multi_platform_whitelistv4&utm_relevant_index=1
> 写起来好麻烦，写的头昏脑涨🤦‍♂️，要是有可以优化的地方记得提醒，因为图的题很少，所以验证起来比较麻烦
> 不过用矩阵转化的方式还是可以得到我们想要的图的
