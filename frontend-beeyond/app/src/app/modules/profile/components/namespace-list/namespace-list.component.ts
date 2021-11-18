import { NestedTreeControl } from '@angular/cdk/tree';
import { Component, OnInit } from '@angular/core';
import { MatTreeNestedDataSource } from '@angular/material/tree';
import { BackendApiService } from 'src/app/core/services/backend-api.service';

interface NamespaceNode {
  name: string;
  children?: NamespaceNode[];
}

@Component({
  selector: 'app-namespace-list',
  templateUrl: './namespace-list.component.html',
  styleUrls: ['./namespace-list.component.scss']
})
export class NamespaceListComponent implements OnInit {
  treeControl = new NestedTreeControl<NamespaceNode>(node => node.children);
  dataSource = new MatTreeNestedDataSource<NamespaceNode>();

  constructor(private backendApiService: BackendApiService) {}

  ngOnInit(): void {
    this.backendApiService.getAllNamespaces().subscribe(namespaces => {
      const tree: NamespaceNode[] = [];

      for (const namespace of namespaces) {
        tree.push({
          name: namespace.namespace,
          children: namespace.users.map(user => ({ name: user.name }))
        });
      }

      this.dataSource.data = tree;
    });
  }

  hasChild = (_, node: NamespaceNode) => !!node.children && node.children.length > 0;
}
