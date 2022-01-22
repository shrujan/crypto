import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { DashboardService } from '../services/dashboard/dashboard.service';

@Component({
  selector: 'app-input-form',
  templateUrl: './input-form.component.html',
  styleUrls: ['./input-form.component.scss']
})
export class InputFormComponent implements OnInit {
  userList: any;
  coinList: any;
  selected: string = '';
  cryptoForm: FormGroup;

  constructor(
    private dashboardService: DashboardService,
  ) {
      this.dashboardService.getUsers().subscribe(userList => {
        this.userList = userList;
      });
      this.dashboardService.getCoins();
      this.initializeForm();
   }

  ngOnInit() {
    this.dashboardService.coinList$.subscribe(coins => this.coinList = coins)
  }

  initializeForm() {
    this.cryptoForm = new FormGroup({
      buySell: new FormControl(),
      coinName: new FormControl(''),
      purchaseDate: new FormControl(),
      transactionPrice: new FormControl(),
      quantity: new FormControl(),
      totalAmount: new FormControl(),
      userName: new FormControl(''),

    })
  }

  saveDetails(): void {
    const cryptoData = this.cryptoForm.value;
    console.log(cryptoData)
    const param = {
      "buySell":          cryptoData.buySell,
      "coinName":         (cryptoData.coinName).toLowerCase(),
      "purchaseDate":     cryptoData.purchaseDate || 'N/A',
      "quantity":         cryptoData.quantity,
      "totalAmount":      cryptoData.totalAmount || (parseInt(cryptoData.quantity) * parseInt(cryptoData.transactionPrice)),
      "transactionPrice": cryptoData.transactionPrice,
      "userName":         (cryptoData.userName.UserName).toLowerCase(),
    }

    this.dashboardService.saveCryptoInfo(param).subscribe(res => {
      console.log("success")
    }) 
  }

}
