һ user�յ�check����֤����
1. request.value = check.value

2. request.token = check.token

3. request.from = check.from

4. request.to = check.to

5. request.opAddr = check.opAddr

6. signer(check) = check.opAddr

7. check.nonce >= contract.nonce[to]

8. check.nonce �Ϸ��Լ��

  ���user�ڱ����Ѵ���һ��check������to��nonce�������յ���chekck��ͬ�����checkΪ�쳣֧Ʊ��operator������֧Ʊ���������ͬһ��toAddr���ܷ�������nonce��ͬ��֧Ʊ��������������������Ҫôoperator����bug��Ҫô����αװ��operator��������֧Ʊ��������operator˽Կй¶����������������յ���check�������û�������operator�쳣����ؾ��档



�� provider����paycheckʱ��֤����

1. signer(check) = paycheck.operator

2. signer(paycheck) = paycheck.from

3. paycheck.payvalue >= 0

4. paycheck.payvalue <= paycheck.value

5. check.nonce >= contract.nonce

6. check.nonce�ظ�ֵ���

   ��Ϊ����һ����check��֤�У�user�Ѿ���֧Ʊ��nonce�ظ��Խ������Ų飬����������user�ǲ����ܷ����ظ�nonce��֧Ʊ�ġ��������ж����user�������operator��sk��Ȼ���check�������޸��Ժ�ʹ��skģ����operator��ǩ����Ȼ���ټ����Լ���userǩ�����ſ��ܳ���provider���յ��ظ�nonce�������

   ������ͬ�ϣ�������ͬnonce��checkֱ�Ӷ���������provider�����user�쳣�ľ��档

