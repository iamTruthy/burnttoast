So working with Go internals....

in the code example:

a := []string{"sunhine", "bumblebee", "icecream", "vanillaice"}
b := a
fmt.Println(a)
fmt.Println(b)

[sunhine bumblebee icecream vanillaice] --a
[sunhine bumblebee icecream vanillaice] --b

we see that A is a []string and so is B which also is and assigned to the same elements as A
b := a

[sunhine bumblebee icecream vanillaice] --a
[sunhine bumblebee icecream vanillaice] --b

------------------------------------------------------------------------------------

This means, both A and B are both pointing to the same underlying array. Therefore whatever change is made to A appears in B. We can see here thats as "banana milkshake" is assigned the index position of [0] in A (a[0] = "banana milkshake") the same change also happens in B as well.

a[0] = "banana milkshake"
fmt.Println(a)
fmt.Println(b)

[banana milkshake bumblebee icecream vanillaice] --a
[banana milkshake bumblebee icecream vanillaice] --b

Since B is "assigned" with the same variables as A most importantly, why any change in A happens in B  because they point to the same underlying array. And the fix for this is having them point to diffrent arrays.

----------------------------------------------------------------------------------------

In the code below, i and x are both [6]int (they both store 6 characters) but i is defined and x is not. we use a COPY function 
**"copy(x, i)"** to COPY the elements of i into x. WE DO NOT ASSIGN the same of x to i because this will have then pointing to the same underlying array. 

i := []int{2, 4, 6, 8, 16, 14}
x := make([]int, 6)
copy(x, i)
fmt.Println(i)
fmt.Println(x)


This way, we have the same elements in i and x both pointing to different underlying arrays. so when 16 is assigned the index of "0"
in the slice, the change only happens in i and not x even though, they both have same elements simply because x is not equal i and assigned the same values as i

i[0] = 16
fmt.Println(i)
fmt.Println(x)
