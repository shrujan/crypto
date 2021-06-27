import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  // baseUrl = 'https://api.coingecko.com/api/v3';
  baseUrl = 'http://localhost:8081';

  constructor(private http: HttpClient) { }

  getCoinsInfoInr() {
    // const url = this.baseUrl + '/coins/markets';
    // const param = {
    //   vs_currency: 'inr',
    //   order: 'market_cap_desc',
    //   sparkline: false
    // }
    // return this.http.get(url, { params: param });
    let url = `${ this.baseUrl }/getMarketInfo`;
    return this.http.get(url);
  }

  sendMail() {
    let url = `${ this.baseUrl }/getMarketInfows`;
    return this.http.get(url);
  }
}
