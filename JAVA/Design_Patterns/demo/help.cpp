#include <iostream>
#include<bits/stdc++.h>


using namespace std;

int getMinimumMoves(vector<vector<int>>maze,int k){

    int n = maze.size() , m = maze[0].size();
    if(maze[0][0] == 1 || maze[n-1][m-1]) return -1;
    queue<pair<int,int>>q;

    int level = 0;
    q.push({0,0});
    maze[0][0] = 1;

    while(!q.empty()){

        for(int i=q.size();i>0;i--)
        {
            pair<int,int>p = q.front();
            q.pop();

            int u = p.first , v = p.second;
            if(u==n-1 && v==m-1) return level+1;
            for(int i=1;i<=k;i++){
                if(u+i < n && maze[u+i][v] == 0){
                    q.push({u+i,v}) ; 
                    maze[u+i][v] = 1;
                }
                if(u-i >=0 && maze[u-i][v] == 0){
                    q.push({u-i,v}) ; 
                    maze[u-i][v] = 1;
                }
                if(v+i < m && maze[u][v+i] == 0){
                    q.push({u,v+i}) ; 
                    maze[u][v+i] = 1;
                }if(v-i >=0 && maze[u][v-i] == 0){
                    q.push({u,v-i}) ; 
                    maze[u][v-i] = 1;
                }
            }
        }
        level++;
    }
    return -1;
}



int main() {
    cout<<"Hello World!";
}
