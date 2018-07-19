import { Component, OnInit } from '@angular/core';
import { VariableService } from '../../services/variable.service';
import { Variable } from '../../models/variable';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-variable',
  templateUrl: './variable-view.component.html',
  styleUrls: []
})
export class VariableComponent implements OnInit {

  variables: Variable[];
  variable: Variable;

  constructor(private variableService: VariableService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.variableService.getVariables().then(variables => this.variables = variables);
  }

  newVariable(): void {

    this.router.navigate(['/variable/new']).then(() => null);
    this.globals.currentModule = 'Variable';
  }

  editar(variable: Variable): void {
    this.variable = variable;
    this.router.navigate(['/variable/edit', this.variable._id ]);
  }

  borrar(variable: Variable): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar variable?',
      accept: () => {
        this.variableService.delete(variable._id)
          .then(response => this.variableService.getVariables().then(variables => this.variables = variables));
      }
    });
  }
}