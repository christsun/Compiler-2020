/*
Test Package: Sema_Local_Preview
Test Target: Classes
Author: Admin
Time: 2019-10-20
Verdict: Success
Origin Package: Semantic Extended
*/
A b;
int c = b.a;

int main() {
}

class A{
    int a;
    void fun() {
    	 a = c;
    }
}

