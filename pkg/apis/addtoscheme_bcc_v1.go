/*
 * @Description:
 * @Autor: gaoshiyao
 * @Date: 2021-08-06 16:26:24
 * @LastEditors: gaoshiyao
 * @LastEditTime: 2021-09-02 19:32:30
 */
package apis

import (
	v1 "github.com/open-beagle/awecloud-sysoperator-sdk/pkg/apis/bcc/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme, v1.SchemeGatewayBuilder.AddToScheme)
}
