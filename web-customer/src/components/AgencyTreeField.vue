<template>
  <el-card class="box-card card-info" style="margin-bottom: 1rem">
    <template #header>
      <div class="card-header">
        <h2 class="card-title">{{ title }}</h2>
      </div>
    </template>
    <div class="text item" style="margin-bottom: 0.5rem">
      <el-tree
        :props="defaultProps"
        :load="loadNode"
        lazy
        @node-click="(node) => $emit('onNodeClick', node)"
      />
    </div>
  </el-card>
</template>

<script>
export default {
  name: "AgencyTreeField",
  props: ["title"],
  emits: ["onNodeClick"],
};
</script>

<script setup>
import { getAgenciesTreeForField, getDocumentAgencyList } from "@/api/agency";

const defaultProps = {
  label: "name",
  children: "children",
  isLeaf: "leaf",
};

const loadNode = async (node, resolve) => {
  if (node.level === 0) {
    const response = await getDocumentAgencyList();
    if (response.code === 0) {
      return resolve(response.data.list);
    }

    return resolve([]);
  }

  const agencyId = node.data.ID;

  const response = await getAgenciesTreeForField({ id: agencyId });
  if (response.code === 0) {
    if (
      !response.data.nodes ||
      !Array.isArray(response.data.nodes) ||
      response.data.nodes.length === 0
    )
      return resolve([{ name: "Không có dữ liệu", leaf: true }]);

    const data = response.data.nodes.map((n) => {
      return {
        ...n,
        leaf: true,
        name: `${n.field.name} - (${n.count})`,
        agencyId: n.agencyId,
        fieldId: n.fieldId,
      };
    });

    return resolve(data);
  }

  return resolve([{ name: "Không có dữ liệu", leaf: true }]);
};
</script>

<style scoped>
.card-info {
  text-align: left;
}
.card-title {
  text-align: left;
  font-size: 1rem;
}
.ml-2 {
  margin-left: 0.3rem;
}
</style>
