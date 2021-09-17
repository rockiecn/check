一 user收到check的验证流程
1. request.value = check.value

2. request.token = check.token

3. request.from = check.from

4. request.to = check.to

5. request.opAddr = check.opAddr

6. signer(check) = check.opAddr

7. check.nonce >= contract.nonce[to]

8. check.nonce 合法性检查

  如果user在本地已存在一个check，它的to和nonce属性与收到的chekck相同，则此check为异常支票（operator发出的支票，正常情况同一个toAddr不能发出两个nonce相同的支票，如果出现这种情况，则要么operator出的bug，要么有人伪装成operator发出恶意支票，可能有operator私钥泄露的情况。），则丢弃收到的check，并向用户发出此operator异常的相关警告。



二 provider接收paycheck时验证内容

1. signer(check) = paycheck.operator

2. signer(paycheck) = paycheck.from

3. paycheck.payvalue >= 0

4. paycheck.payvalue <= paycheck.value

5. check.nonce >= contract.nonce

6. check.nonce重复值检查

   因为在上一步的check验证中，user已经对支票的nonce重复性进行了排查，所以正常的user是不可能发出重复nonce的支票的。除非是有恶意的user，获得了operator的sk，然后对check的内容修改以后，使用sk模拟了operator的签名，然后再加上自己的user签名，才可能出现provider接收到重复nonce的情况。

   处理方法同上，发现相同nonce的check直接丢弃，并向provider报告此user异常的警告。

