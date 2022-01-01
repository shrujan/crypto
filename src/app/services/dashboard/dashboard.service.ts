import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs';

export interface User {
  UserName: string,
  Email: string
} 

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  // baseUrl = 'https://api.coingecko.com/api/v3';
  baseUrl        = 'http://localhost:8081';
  coinListInr$ = new BehaviorSubject<any>(undefined);


  constructor(private http: HttpClient) { }

  getCoinsInfoInr(): void {
    let url = `${ this.baseUrl }/getMarketInfo`;
    this.http.get(url).subscribe(coinList => {
      this.coinListInr$.next(coinList)
    });
  }

  getUsers() {
    let url = `${ this.baseUrl }/getUsers`;
    return this.http.get(url);
  }

  saveCryptoInfo(data: any) {
    let url = `${ this.baseUrl }/savePurchaseInfo`;
    return this.http.post(url, data);
  }

  sendMail() {
    let url = `${ this.baseUrl }/getMarketInfows`;
    return this.http.get(url);
  }
}
