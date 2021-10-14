//go:generate go run -tags generate ../../generate/tags/main.go -GetTag=yes -ListTags=yes -ListTagsOp=DescribeTags -ListTagsInFiltIDName=auto-scaling-group -ServiceTagsSlice=yes -TagOp=CreateOrUpdateTags -TagResTypeElem=ResourceType -TagType2=TagDescription -TagTypeAddBoolElem=PropagateAtLaunch -TagTypeIDElem=ResourceId -UntagOp=DeleteTags -UntagInNeedTagType=yes -UntagInTagsElem=Tags -UpdateTags=yes
// ONLY generate directives and package declaration! Do not add anything else to this file.

package autoscaling
