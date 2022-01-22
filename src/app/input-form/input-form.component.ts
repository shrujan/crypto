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
      purchasePrice: new FormControl(),
      quantity: new FormControl(),
      totalAmount: new FormControl(),
      userName: new FormControl(''),

    })
  }

  saveDetails(): void {
    const cryptoData = this.cryptoForm.value;
    console.log(cryptoData)
    const param = {
      "buySell": cryptoData.buySell,
      "userName": (cryptoData.userName.UserName).toLowerCase(),
      "coinName": (cryptoData.coinName).toLowerCase(),
      "quantity": cryptoData.quantity,
      "purchasePrice": cryptoData.purchasePrice,
      "purchaseDate": cryptoData.purchaseDate || 'N/A',
      "totalAmount": cryptoData.totalAmount || (parseInt(cryptoData.quantity) * parseInt(cryptoData.purchasePrice))
    }

    this.dashboardService.saveCryptoInfo(param).subscribe(res => {
      console.log("success")
    }) 
  }

}
