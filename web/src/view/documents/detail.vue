<template>
  <div>Detail document: {{ document?.title }}</div>
</template>

<script>
export default {
  name: 'DocumentDetail',
}
</script>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { findDocuments } from '../../api/documents'

const route = useRoute()

const document = ref(null)

const searchInfo = ref({
  ID: Number(route.params.document_id)
})

watch(() => route.params.document_id, (id) => {
  searchInfo.value.ID = Number(id)
  getDocument()
})

const getDocument = async() => {
  if (!Number.isInteger(searchInfo.value.ID)) { return }

  const data = await findDocuments(searchInfo.value)
}

getDocument()

</script>
