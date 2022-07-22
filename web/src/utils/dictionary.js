import { useDictionaryStore } from '@/pinia/modules/dictionary'
// Get dictionary method using example getDict('sex').then(res) or const res = await getDict('sex') under async function
export const getDict = async(type) => {
  const dictionaryStore = useDictionaryStore()
  await dictionaryStore.getDictionary(type)
  return dictionaryStore.dictionaryMap[type]
}
