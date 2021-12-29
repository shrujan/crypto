import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface User {
  UserName: string,
  Email: string
} 

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  // baseUrl = 'https://api.coingecko.com/api/v3';
  baseUrl = 'http://localhost:8081';

  constructor(private http: HttpClient) { }

  getCoinsInfoInr() {
    let url = `${ this.baseUrl }/getMarketInfo`;
    return this.http.get(url);
  }

  getUsers() {
    let url = `${ this.baseUrl }/getUsers`;
    return this.http.get(url);
  }

  sendMail() {
    let url = `${ this.baseUrl }/getMarketInfows`;
    return this.http.get(url);
  }
}
