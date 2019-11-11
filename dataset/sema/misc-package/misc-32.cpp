/*
Test Package: Sema_Local_Preview
Test Target: Misc
Author: Pikachu
Time: 2019-11-11
Verdict: Fail
Comment: If condition can only be bool(Line 7), -> not defined(Line 51), Duplicated variables(Line 147)
Origin Package: Semantic Pretest
*/
// A changed solution for 1630 on SJTU Online Judge 
int maxr = 25;
int maxc = 20005;
int INF = -1 >> 1;

int mmin(int a, int b) { 
    if (a + b) {
        return a;
    } else {
        return b;
    }
}   

int mmax(int a, int b) { 
    if (a < b) {
        return b;
    } else {
        return a;
    }
}

class lnode {
    int mint;
    int maxt;
    int sumt;
    int ptag;
    int mtag;
    lnode() {
        mint = 0;
        maxt = 0;
        sumt = 0;
        ptag = 0;
        mtag = 1;
    }
}

lnode[][] ltree = new lnode[maxr][4 * maxc];
int[][] matrix = new int[maxr][maxc];
int r;
int  c; 
int  m; 

void pushdown(int i, int x, int l, int r) {
    int ptag = ltree[i][x].ptag;
    int mtag = ltree[i][x].mtag;
    int mid = (l + r) / 2;
    ltree[i][2 * x].mtag = ltree[i][2 * x].mtag * mtag;
    ltree[i][2 * x].ptag = ltree[i][2 * x].ptag * mtag + ptag;
    ltree[i][2 * x + 1].mtag = ltree[i][2 * x + 1].mtag * mtag;
    ltree[i][2 * x + 1].ptag = ltree[i][2 * x + 1].ptag * mtag + ptag;
    ltree[i][2 * x].mint = ltree[i][2 * x].mint * mtag + ptag;
    ltree[i][2 * x + 1]->mint = ltree[i][2 * x + 1].mint * mtag + ptag;
    ltree[i][2 * x].maxt = ltree[i][2 * x].maxt * mtag + ptag;
    ltree[i][2 * x + 1].maxt = ltree[i][2 * x + 1].maxt * mtag + ptag;
    ltree[i][2 * x].sumt = ltree[i][2 * x].sumt * mtag + ptag * (mid - l + 1);
    ltree[i][2 * x + 1].sumt =
        ltree[i][2 * x + 1].sumt * mtag + ptag * (r - mid);
    ptag = 0;
    mtag = 1;
}

void modifyadd(int i, int x, int l, int r, int ql, int qr, int v) {
    if (l >= ql && r <= qr) {
        ltree[i][x].sumt = ltree[i][x].sumt + (r - l + 1) * v;
        ltree[i][x].ptag = ltree[i][x].ptag + v;
        ltree[i][x].mint = ltree[i][x].mint + v;
        ltree[i][x].maxt = ltree[i][x].maxt + v;
        return;
    }
    int mid = (l + r) / 2;
    pushdown(i, x, l, r);
    if (ql <= mid) modifyadd(i, 2 * x, l, mid, ql, qr, v);
    if (qr >= mid + 1) modifyadd(i, 2 * x + 1, mid + 1, r, ql, qr, v);
    ltree[i][x].sumt = ltree[i][2 * x].sumt + ltree[i][2 * x + 1].sumt;
    ltree[i][x].mint = mmin(ltree[i][2 * x].mint, ltree[i][2 * x + 1].mint);
    ltree[i][x].maxt = mmax(ltree[i][2 * x].maxt, ltree[i][2 * x + 1].maxt);
}

void modifymulti(int i, int x, int l, int r, int ql, int qr, int v) {
    if (l >= ql && r <= qr) {
        ltree[i][x].sumt = ltree[i][x].sumt *v;
        ltree[i][x].mint = ltree[i][x].mint *v;
        ltree[i][x].maxt = ltree[i][x].maxt *v;
        ltree[i][x].mtag = ltree[i][x].mtag *v;
        ltree[i][x].ptag = ltree[i][x].ptag *v;
        return;
    }
    int mid = (l + r) / 2;
    pushdown(i, x, l, r);
    if (ql <= mid) modifymulti(i, 2 * x, l, mid, ql, qr, v);
    if (qr >= mid + 1) modifymulti(i, 2 * x + 1, mid + 1, r, ql, qr, v);
    ltree[i][x].sumt = ltree[i][2 * x].sumt + ltree[i][2 * x + 1].sumt;
    ltree[i][x].mint = mmin(ltree[i][2 * x].mint, ltree[i][2 * x + 1].mint);
    ltree[i][x].maxt = mmax(ltree[i][2 * x].maxt, ltree[i][2 * x + 1].maxt);
}

int asksum(int i, int x, int l, int r, int ql, int qr) {
    int sum = 0;
    if (l >= ql && r <= qr) {
        sum = sum + ltree[i][x].sumt;
        return sum;
    }
    pushdown(i, x, l, r);
    int mid = (l + r) / 2;
    if (ql <= mid) sum = sum + asksum(i, 2 * x, l, mid, ql, qr);
    if (qr >= mid + 1) sum = sum + asksum(i, 2 * x + 1, mid + 1, r, ql, qr);
    return sum;
}

int askmin(int i, int x, int l, int r, int ql, int qr) {
    if (l >= ql && r <= qr) {
        return ltree[i][x].mint;
    }
    pushdown(i, x, l, r);
    int mid = (l + r) / 2;
    int ansmin = INF;
    if (ql <= mid) ansmin = mmin(askmin(i, 2 * x, l, mid, ql, qr), ansmin);
    if (qr >= mid + 1)
        ansmin = mmin(askmin(i, 2 * x + 1, mid + 1, r, ql, qr), ansmin);
    return ansmin;
}

int askmax(int i, int x, int l, int r, int ql, int qr) {
    if (l >= ql && r <= qr) {
        return ltree[i][x].maxt;
    }
    pushdown(i, x, l, r);
    int mid = (l + r) / 2;
    int ansmax = 0;
    if (ql <= mid) ansmax = mmax(askmax(i, 2 * x, l, mid, ql, qr), ansmax);
    if (qr >= mid + 1)
        ansmax = mmax(askmax(i, 2 * x + 1, mid + 1, r, ql, qr), ansmax);
    return ansmax;
}

int main() {
    r = getInt();
    c = getInt();
    m = getInt();

    int op;
    int  x1; 
    int  y1; 
    int  x2; 
    int  y2, v; 
    int i; 
    int y2;
    for (i = 1; i <= m; ++i) {
        op = getInt();
        x1 = getInt();
        y1 = getInt();
        x2 = getInt();
        y2 = getInt();
        if (op == 1) {
            v = getInt();
            int i;
            for (i = x1; i <= x2; ++i) {
                modifyadd(i, 1, 1, c, y1, y2, v);
            }
        } else if (op == 2) {
            v = getInt();
            int i;
            for (i = x1; i <= x2; ++i) {
                modifymulti(i, 1, 1, c, y1, y2, 0);
                modifyadd(i, 1, 1, c, y1, y2, v);
            }
        } else if (op == 3) {
            int ans = 0;
            int i;
            for (i = x1; i <= x2; ++i) {
                ans = ans + asksum(i, 1, 1, c, y1, y2);
            }
            println(toString(ans) + "\n");
        } else {
            int ansmin = INF;
            int  ansmax = 0; 
            int i;
            for (i = x1; i <= x2; ++i) {
                ansmin = mmin(ansmin, askmin(i, 1, 1, c, y1, y2));
                ansmax = mmax(ansmax, askmax(i, 1, 1, c, y1, y2));
            }
            println(toString(ansmin) + " " + toString(ansmax) + "\n");
        }
    }
    return 0;
}