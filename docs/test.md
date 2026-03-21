# 链表

链表的场景常见操作：插入、删除、查找、遍历

链表分带虚拟头节点(推荐)和不带虚拟头节点(头指针)两种实现，在涉及头节点的操作，为统一可使用虚拟头节点，尾节点为空指针，最后结果需返回虚拟头节点.Next

**巧思**：快慢指针找中间节点，翻转单链表(整体或部分)

## **单链表**

注意增和删改查边界问题，手动添加虚拟头节点便于解决问题

**设计链表**

[707. 设计链表 - 力扣（LeetCode）](https://leetcode.cn/problems/design-linked-list/description/)

```go
// 节点结构体
type Node struct{
    val int
    next *Node
}
// 单链表结构体
type MyLinkedList struct {
    Head *Node
    Size int
}

func Constructor() MyLinkedList {
    return MyLinkedList{Head: &Node{},Size: 0}
}
// 获取指定下标的节点
func (this *MyLinkedList) goToIndex(index int) *Node{
    count := -1
    node := this.Head
    for count != index {
        node = node.next
        count ++
    }
    return node
}

func (this *MyLinkedList) Get(index int) int {
    if index >= this.Size || index < 0{
        return -1
    }
    return this.goToIndex(index).val
}

func (this *MyLinkedList) AddAtHead(val int)  {
    this.Head.next = &Node{val: val,next: this.Head.next}
    this.Size++
}

func (this *MyLinkedList) AddAtTail(val int)  {
    node := this.goToIndex(this.Size-1)
    node.next = &Node{val: val}
    this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.Size || index < 0{
        return
    }
    node := this.goToIndex(index-1)
    newNode := &Node{val: val, next: node.next}
    node.next = newNode
    this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.Size || index < 0{
        return
    }
    node := this.goToIndex(index-1)
    node.next = node.next.next
    this.Size--
}
```



### 找链表节点

#### 找单链表中间节点

链表节点数为偶数时会出现找前中点和后中点的问题

推荐slow停止在前中点，fast停止在末尾节点的方式

[876. 链表的中间结点 - 力扣（LeetCode）](https://leetcode.cn/problems/middle-of-the-linked-list/submissions/703913742/)

```go
// slow停在中间节点，fast停在nil
func middleNode(head *ListNode) *ListNode {
    dummyNode := &ListNode{Next: head}	// 虚拟头节点
    fast,slow := dummyNode, dummyNode
    // 找目标节点
    for fast!=nil{
        fast=fast.Next
        if fast!=nil{
            fast=fast.Next
        }
        slow=slow.Next
    }
    return slow
}

// slow停止在前中点，fast停止在最后一个节点
func middleNode(head *ListNode) *ListNode {
    dummyNode := &ListNode{Next: head}	// 虚拟头节点
    fast,slow := dummyNode, dummyNode
    // 找目标节点
    for fast.Next!=nil{
        fast=fast.Next
        if fast.Next!=nil{
            fast=fast.Next
            slow=slow.Next
        }
    }
    return slow.Next
}
```

#### **找链表倒数第N个节点**

指针移动条件会影响`slow`和`fast`最后停留的位置：

`fast!=nil`时`fast`停留在nil，`slow`停留在目标节点；

`fast.Next!=nil`时`fast`停留在最后一个节点，`slow`停留在目标节点的前一个节点

[LCR 140. 训练计划 II - 力扣（LeetCode）](https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)

```go
// slow停在倒数第N个节点，fast停在nil
func trainingPlan(head *ListNode, cnt int) *ListNode {
    dummyNode := &ListNode{Next: head}
    fast, slow := dummyNode, dummyNode
    for i:=0;i<cnt;i++{
        fast=fast.Next
    }
    for fast!=nil{
        fast = fast.Next
        slow = slow.Next
    }
    return slow
}

// slow停在倒数第N个节点前，fast停在末节点
func trainingPlan(head *ListNode, cnt int) *ListNode {
    dummyNode := &ListNode{Next: head}
    fast, slow := dummyNode, dummyNode
    for i:=0;i<cnt;i++{
        fast=fast.Next
    }
    for fast.Next!=nil{
        fast = fast.Next
        slow = slow.Next
    }
    return slow.Next
}
```

**旋转链表**

类似[轮转数组](#轮转数组)

[61. 旋转链表 - 力扣（LeetCode）](https://leetcode.cn/problems/rotate-list/submissions/703220512/)

```go
func rotateRight(head *ListNode, k int) *ListNode {
    if head==nil{
        return nil
    }
    // 需要先遍历链表长度再取余
    length:=0
    list:=head
    for list!=nil{
        length++
        list=list.Next
    }
    // 本质是快慢指针找到链表倒数第K%length+1个节点
    slow, fast := head, head
    for i:=0;i<k%length;i++{
        fast=fast.Next
    }
    for fast.Next!=nil{
        fast=fast.Next
        slow=slow.Next
    }
    fast.Next=head
    head=slow.Next
    slow.Next=nil
    return head
}
```



### 删除链表节点

#### 删除链表中的节点

无法获取上一个节点且可修改节点的值：复制下一个节点的值并删除下一个节点

[237. 删除链表中的节点 - 力扣（LeetCode）](https://leetcode.cn/problems/delete-node-in-a-linked-list/submissions/683222115/)

```go
func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}
```

#### **删除链表倒数第N个节点**

本质是找倒数第N+1个节点

[19. 删除链表的倒数第 N 个结点 - 力扣（LeetCode）](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/)

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 写操作建议使用虚拟头节点
    dummyNode:=&ListNode{Next: head}
    fast,slow := dummyNode,dummyNode
    // 本质是找倒数第N+1个节点
    for i:=0;i<n;i++{
        fast=fast.Next
    }
    for fast.Next!=nil{
        fast=fast.Next
        slow=slow.Next
    }
    slow.Next = slow.Next.Next
    return dummyNode.Next
}
```

#### **移除链表元素**

[203. 移除链表元素 - 力扣（LeetCode）](https://leetcode.cn/problems/remove-linked-list-elements/description/)

需要注意只在不用删除时才手动移动，删除逻辑已经实现移动了

```go
func removeElements(head *ListNode, val int) *ListNode {
    // var node *ListNode 只声明不分配内存会报空指针异常
    node := &ListNode{Next: head}   // 也可使用new()
    head = node // 复用
    for head.Next != nil{
        if head.Next.Val == val{
            head.Next = head.Next.Next
        }else{// 需要注意只在不用删除时才移动，删除时已经移动了
            head = head.Next
        }
    }
    return node.Next
}
```

#### 删除排序列表重复元素

[移除元素](#移除元素)

[83. 删除排序链表中的重复元素 - 力扣（LeetCode）](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/)

去重，需要注意只在不用删除时才手动移动，删除逻辑已经实现移动了

```go
func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    dummyNode:=&ListNode{Next: head}
    // for+if
    for head.Next!=nil{
        if head.Val==head.Next.Val{
            head.Next=head.Next.Next
        }else{
            // 需要注意只在不用删除时才手动移动，删除逻辑已经实现移动了
            head=head.Next
        }
    }
    return dummyNode.Next
}
```

[82. 删除排序链表中的重复元素 II - 力扣（LeetCode）](https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/)

不是去重而是全去，删除链表节点一般都需要pre记录前节点

使用一个flag标识当前元素是不是最后一个重复元素

```go
func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    dummyNode:=&ListNode{Next: head}
    pre:=dummyNode
    flag:=false
    for flag || (head!=nil&&head.Next!=nil){
        // 当前节点与下一节点有重复值：移除下一节点
        if head.Next!=nil && head.Val==head.Next.Val{
            flag=true
            head.Next=head.Next.Next
        }else{
            // 移除全部相同元素
            if flag{
                pre.Next=head.Next
                flag=false
            }else{
                // 正常移动
                pre=head
            }
            head=head.Next
        }
    }
    return dummyNode.Next
}
```





### **单链表翻转**

#### 翻转整个链表

[206. 反转链表 - 力扣（LeetCode）](https://leetcode.cn/problems/reverse-linked-list/submissions/661146417/?envType=study-plan-v2&envId=top-100-liked)

双指针解法：本质是三指针

需注意这一性质：反转后，pre指向原节点末尾节点，cur指向下一节点

```go
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    cur:=head
    for cur!=nil{
        // 指针修改顺序画图易于理解
        next:=cur.Next  // 需要next暂存下一节点
        cur.Next=pre
        pre=cur
        cur=next
    }
    // 翻转链表的一大特点是pre最后停留在末节点，cur指向下一节点
    return pre
}
```

递归解法

```go
func Reverse(cur, pre *Node){
    if cur == nil{
		return pre
    }
    next := cur.next
	cur.next = pre
    Reverse(next, cur)
}

Reverse(MyLinkedList.Head, nil)
```

#### 翻转链表一部分

先移动node到cur.Pre，再进行链表翻转

[92. 反转链表 II - 力扣（LeetCode）](https://leetcode.cn/problems/reverse-linked-list-ii/)

```go
需注意，部分翻转后：
pre停留在翻转链表的首节点节点
cur停留在后半段链表的首节点
left停留在上半段链表的末节点且left.Next指向翻转链表的末节点
```

<img src="C:/Users/RaY/D/Note/TyporeNote/后端开发/算法/算法.assets/image-20251111150303757.png" alt="image-20251111150303757" style="zoom:33%;" />

```go
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummyNode:=&ListNode{Next: head}
    node:=dummyNode
    // 先移动到翻转部分的前一节点
    for i:=0;i<left-1;i++{
        node=node.Next
    }
    var pre *ListNode
    cur:=node.Next
    // 根据需翻转的节点执行次数
    for i:=0;i<right-left+1;i++{
        // 同翻转链表
        next:=cur.Next
        cur.Next=pre
        pre=cur
        cur=next
    }
    // 翻转部分的前一节点用来重接链表
    // 关键一步,node.Next.Next指向翻转后的尾节点,只需指向cur
    node.Next.Next=cur 
    node.Next=pre
    return dummyNode.Next
}
```

**回文单链表**

找中间节点+翻转单链表

[234. 回文链表 - 力扣（LeetCode）](https://leetcode.cn/problems/palindrome-linked-list/?envType=study-plan-v2&envId=top-100-liked)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 回文单链表：找中间节点+翻转单链表
func isPalindrome(head *ListNode) bool {
    // 快慢指针找中间节点
    fast, slow := head, head
    for fast.Next!=nil{
        fast=fast.Next
        if fast!=nil&&fast.Next!=nil{
            fast=fast.Next
        }
        slow=slow.Next
    }

    // 翻转单链表
    var pre *ListNode
    cur := slow.Next
    for cur!=nil{
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    // cur最后是nil，pre为翻转后的首节点
    // 将判断回文链表的问题转换为判断两个链表是否对应位置值相同
    // List1在奇数链表的情况下会比List2多一个节点，但不影响
    for head!=nil&&pre!=nil{
        if head.Val!=pre.Val{
            return false
        }
        head=head.Next
        pre=pre.Next
    }
    return true
}
```

#### K个一组翻转链表

[25. K 个一组翻转链表 - 力扣（LeetCode）](https://leetcode.cn/problems/reverse-nodes-in-k-group/description/)

链表反转后需要移动node到当前组的最后一个元素(pre)以便翻转下一组

<img src="C:\Users\RaY\AppData\Roaming\Typora\typora-user-images\image-20260314213257989.png" alt="image-20260314213257989" style="zoom: 33%;" /> 

```go
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummyNode:=&ListNode{Next: head}
    
    // 记录链表长度
    node:=dummyNode.Next
    length:=0
    for node!=nil{
        length++
        node=node.Next
    }
    
    left:=dummyNode
    var pre *ListNode
    cur:=left.Next
    // 按k个一组翻转
    for ;length>=k;length-=k{
        // k个一组
        for i:=0;i<k;i++{
            next:=cur.Next
            cur.Next=pre
            pre=cur
            cur=next
        }
		// 画图便全部理解了
        
        // 需记录next已移动node到下一组的前一节点
        right:=left.Next

        // 同翻转部分链表
        left.Next.Next=cur
        left.Next=pre

        // 移动到下一组的前一节点
        left=right
    }

    return dummyNode.Next
}
```

### 合并链表

#### 合并两个升序链表

[合并两个有序数组](#合并两个有序数组)

同时不为空且处理其中一个不为空

[21. 合并两个有序链表 - 力扣（LeetCode）](https://leetcode.cn/problems/merge-two-sorted-lists/submissions/704250866/)

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummyNode:=&ListNode{}
    node:=dummyNode
    // 同时要不为nil
    for list1!=nil&&list2!=nil{
        if list1.Val<=list2.Val{
            node.Next=list1
            list1=list1.Next
        }else{
            node.Next=list2
            list2=list2.Next
        }
        node=node.Next
    }
    // 处理其中一个不为nil的
    // 单链表已经串起来，只需要if
    if list1!=nil{
        node.Next=list1
    }
    if list2!=nil{
        node.Next=list2
    }
    return dummyNode.Next
}
```

**链表相加**

本质同合并两个升序链表

[2. 两数相加 - 力扣（LeetCode）](https://leetcode.cn/problems/add-two-numbers/submissions/704296866/?envType=study-plan-v2&envId=top-100-liked)

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyNode:=&ListNode{}
    node:=dummyNode
    
    // 参考合并两个升序链表
    carry:=&ListNode{}
    for l1!=nil||l2!=nil||carry.Val!=0{
        temp:=&ListNode{Val: carry.Val}
        carry.Val=0
        if l1!=nil{
            temp.Val+=l1.Val
            l1=l1.Next
        }
        if l2!=nil{
            temp.Val+=l2.Val
            l2=l2.Next
        }
        if temp.Val/10==1{
            carry.Val=temp.Val/10
            temp.Val%=10
        }
        node.Next=temp
        node=node.Next
    }
    return dummyNode.Next
}
```

#### 重排链表

找链表中点+翻转后半部分+链表合并

[143. 重排链表 - 力扣（LeetCode）](https://leetcode.cn/problems/reorder-list/)

```go
func reorderList(head *ListNode)  {
    dummyNode:=&ListNode{Next: head}
    // 找链表中间节点
    fast, slow:=dummyNode, dummyNode
    for fast!=nil{
        fast=fast.Next
        if fast!=nil{
            fast=fast.Next
        }
        slow=slow.Next
    }

    // 翻转后半部分链表
    var pre *ListNode
    cur:=slow.Next
    slow.Next=nil
    for cur!=nil{
        next:=cur.Next
        cur.Next=pre
        pre=cur
        cur=next
    }

    // 合并两个链表
    list2:=pre
    node:=dummyNode.Next
    for list2!=nil{
        next:=node.Next
        node.Next=list2

        next2:=list2.Next
        list2.Next=next

        list2=next2

        node=next
    }    
}
```

#### 合并K个升序链表

两两合并而不是从左到右两两合并

[23. 合并 K 个升序链表 - 力扣（LeetCode）](https://leetcode.cn/problems/merge-k-sorted-lists/description/)

```go
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists)<1{
        return nil
    }
	// 合并两个链表
    var combine func(list1, list2 *ListNode)*ListNode
    combine = func(list1, list2 *ListNode)*ListNode{
        dummyNode:=&ListNode{}
        node:=dummyNode
        for list1!=nil&&list2!=nil{
            if list1.Val<=list2.Val{
                node.Next=list1
                list1=list1.Next
            }else{
                node.Next=list2
                list2=list2.Next
            }
            node=node.Next
        }
        if list1!=nil{
            node.Next=list1
        }
        if list2!=nil{
            node.Next=list2
        }
        return dummyNode.Next
    }
    
    // 放回原链表
    for len(lists)>1{
        list1:=lists[0]
        lists=lists[1:]
        var list2 *ListNode
        if len(lists)>=1{
            list2=lists[0]
            lists=lists[1:]
        }
        list3 := combine(list1, list2)
        lists=append(lists, list3)
    }
    // 放到新链表
    // 分治:两两合并再两两合并:O(kN)
    // 而不是从左到右合并:O(logk * N)
    for len(lists)>1{
        temp := []*ListNode{}
        // 分治:两两合并再两两合并:O(kN)
        // 而不是从左到右合并:O(logk * N)
        for i:=0; i<len(lists); i+=2{
            list1 := lists[i]
            var list2 *ListNode
            if i+1 < len(lists){
                list2 = lists[i+1]
            }
            result := combine(list1, list2)
            temp =append(temp, result)
        }
        lists = temp
    }

    return lists[0]
}
```

#### 链表排序

单链表排序，先拆分为个体再合并链表，难度取决于空间复杂度的要求

[148. 排序链表 - 力扣（LeetCode）](https://leetcode.cn/problems/sort-list/submissions/704292111/)

简单解法：

空间复杂度O(N)：拆分成节点切片，转为合并K个升序链表

```go
func sortList(head *ListNode) *ListNode {
    if head==nil{
        return nil
    }
    
    dummyNode:=&ListNode{Next: head}
    var lists []*ListNode
    for dummyNode.Next!=nil{
        next:=dummyNode.Next
        dummyNode.Next=next.Next
        next.Next=nil
        lists=append(lists, next)
    }
    
    combine := func(list1, list2 *ListNode)*ListNode{
       list3:=&ListNode{}
       node:=list3
       for list1!=nil&&list2!=nil{
            if list1.Val<=list2.Val{
                node.Next=list1
                list1=list1.Next
            }else{
                node.Next=list2
                list2=list2.Next
            }
            node=node.Next
       }
       if list1!=nil{
        node.Next=list1
       }
       if list2!=nil{
        node.Next=list2
       }
       return list3.Next
    }

    for len(lists)>1{
        list1:=lists[0]
        lists=lists[1:]
        var list2 *ListNode
        if len(lists)>=1{
            list2=lists[0]
            lists=lists[1:]
        }
        list3:=combine(list1,list2)
        lists=append(lists,list3)
    }
    
    return lists[0]
}
```

中等解法：

空间复杂度O(logN)：通过递归找中间节点拆分链表，再转为合并K个升序链表

本质是找单链表中间节点与合并两个链表

[148. 排序链表 - 力扣（LeetCode）](https://leetcode.cn/problems/sort-list/)

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    // 基准条件
    // 单个节点不再拆分
    if head==nil||head.Next==nil{
        return head
    }
    // 找前中点
    dummyNode:=&ListNode{Next: head}
    fast,slow:=dummyNode,dummyNode
    for fast.Next!=nil{
        fast=fast.Next
        if fast.Next!=nil{
            fast=fast.Next
            slow=slow.Next
        }
    }
    // 拆分链表
    left:=dummyNode.Next
    right:=slow.Next
    slow.Next=nil
    // 递归拆分
    left=sortList(left)
    right=sortList(right)
    
    // 从单节点开始合并两个有序链表，自底向上升序合并
    return combine(left, right)
}

func combine(list1, list2 *ListNode)*ListNode{
    list3:=&ListNode{}
    node:=list3
    for list1!=nil&&list2!=nil{
        if list1.Val<=list2.Val{
            node.Next=list1
            list1=list1.Next
        }else{
            node.Next=list2
            list2=list2.Next
        }
        node=node.Next
    }
    if list1!=nil{
    node.Next=list1
    }
    if list2!=nil{
    node.Next=list2
    }
    return list3.Next
}
```

困难解法：

空间复杂度O(1)：

```go
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    // 1. 获取链表长度
    length := 0
    for h := head; h != nil; h = h.Next {
        length++
    }

    // 虚拟头节点
    res := &ListNode{Next: head}
    
    // 2. 自底向上归并，步长 intv: 1, 2, 4, ...
    for intv := 1; intv < length; intv *= 2 {
        pre := res
        h := res.Next
        
        for h != nil {
            // --- 获取左段头 h1 ---
            h1 := h
            i := intv
            // 移动 h 到左段末尾，同时计数实际长度
            for i > 0 && h != nil {
                h = h.Next
                i--
            }
            // 如果 i > 0，说明剩余节点不足 intv，无需再找右段，直接结束本轮
            if i > 0 {
                break
            }

            // --- 获取右段头 h2 ---
            h2 := h
            j := intv
            // 移动 h 到右段末尾，作为下一轮的起始点
            for j > 0 && h != nil {
                h = h.Next
                j--
            }
            
            // 计算左右两段的实际长度
            c1 := intv
            c2 := intv - j // 右段可能不足 intv

            // --- 合并 h1 和 h2 ---
            for c1 > 0 && c2 > 0 {
                if h1.Val <= h2.Val {
                    pre.Next = h1
                    h1 = h1.Next
                    c1--
                } else {
                    pre.Next = h2
                    h2 = h2.Next
                    c2--
                }
                pre = pre.Next
            }

            // 连接剩余部分
            if c1 > 0 {
                pre.Next = h1
            } else {
                pre.Next = h2
            }

            // 将 pre 移动到当前合并段的末尾
            // 注意：此时 pre.Next 指向的是下一段的头 (即当前的 h)
            // 我们需要走 c1 + c2 步到达尾部
            for c1 > 0 || c2 > 0 {
                pre = pre.Next
                if c1 > 0 { c1-- }
                if c2 > 0 { c2-- }
            }
            
            // 连接下一组
            pre.Next = h
        }
    }

    return res.Next
}
```



### **两两交换链表节点**

[24. 两两交换链表中的节点 - 力扣（LeetCode）](https://leetcode.cn/problems/swap-nodes-in-pairs/)

```go
func swapPairs(head *ListNode) *ListNode {
    dummyNode:=&ListNode{Next: head}
    // 初始化在前一节点
    pre:=dummyNode
    fast,slow:=pre,pre
    // 停止在偶数个节点
    for fast.Next!=nil{
        fast=fast.Next
        // 避免奇数个节点
        if fast.Next==nil{
            break
        }
        fast=fast.Next
        slow=slow.Next
        // 交换两个节点
        next:=fast.Next
        fast.Next=slow
        slow.Next=next
        pre.Next=fast

        // 恢复到前一节点
        fast=slow
        pre=slow
    }
    return dummyNode.Next
}
```



### **环形链表**

找到链表第一个成环的节点，哈希辅助可解但空间复杂度O(n)

快慢指针解法

快指针走两步慢指针走一步，快指针走到nil时无环，快慢相遇时即为成环

#### 判断链表是否成环

快慢指针

[141. 环形链表 - 力扣（LeetCode）](https://leetcode.cn/problems/linked-list-cycle/submissions/661777303/?envType=study-plan-v2&envId=top-100-liked)

```go
func hasCycle(head *ListNode) bool {
    if head==nil{
        return false
    }
    dummyNode:=&ListNode{Next: head}
    fast,slow:=dummyNode,dummyNode
    for fast!=nil{
        fast=fast.Next
        if fast!=nil{
            fast=fast.Next
        }
        slow=slow.Next
        if fast==slow{
            return true
        }
    }
    return false
}
```

#### 找到环形链表环入口

确定环入口：快慢相遇位置和头节点各有一个指针同时移动，相遇时为环入口位置

快慢相遇时有环且会在慢指针在环内的第一圈相遇，此时慢指针走了x+y，快指针走了x+ky

[142. 环形链表 II - 力扣（LeetCode）](https://leetcode.cn/problems/linked-list-cycle-ii/submissions/653624284/)

```go
// 快指针走两步慢指针走一步，快指针走到nil时无环，
// 快慢相遇时有环且会在慢指针在环内的第一圈相遇，此时慢指针走了x+y，快指针走了x+ky
// 确定环入口：快慢相遇位置和头节点各有一个指针同时移动，相遇时为环入口位置
func detectCycle(head *ListNode) *ListNode {
    dummyNode:=&ListNode{Next: head}
    fast,slow:=dummyNode,dummyNode
    for fast!=nil&&fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
        // 快慢相遇
        if fast==slow{
            ptr:=dummyNode
            for fast!=ptr{
                fast=fast.Next
                ptr=ptr.Next
            }
            // 环入口
            return ptr
        }
    }
    // 无环
    return nil
}
```

#### **链表相交节点**

[160. 相交链表 - 力扣（LeetCode）](https://leetcode.cn/problems/intersection-of-two-linked-lists/submissions/661143164/?envType=study-plan-v2&envId=top-100-liked)

<img src="C:/Users/RaY/D/Note/TyporeNote/后端开发/算法/算法.assets/image-20250909212133946.png" alt="image-20250909212133946" style="zoom:25%;" /> 

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // a+b+c,走相同的距离,一定在相交点相遇
    ptr1,ptr2:=headA, headB
    for ptr1!=ptr2{
        if ptr1!=nil{
            ptr1=ptr1.Next
        }else{
            ptr1=headB
        }
        if ptr2!=nil{
            ptr2=ptr2.Next
        }else{
            ptr2=headA
        }
    }
    return ptr1
}
```
