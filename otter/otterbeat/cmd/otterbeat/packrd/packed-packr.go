// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "6f832d0391dd840a6d02f9048870911f"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"cb62e23f86addb96ead8c397f3769a97": "1f8b08000000000000ffac54df8b1b451c7f9fbfe29bc972bda8bb9bbb37953da8d4c2419143eb835c63996427c990c9ec303b9b3bd96c11117cb0a80fb51e5af5412c7d39cea7d27205ff994bdab7fe0b32339b5c769313c4dec31ecc7c7f7c7e4d9a8db0cb44d825e910217a2c13a5e160ffc6dd9bfbb73e8c26448592c5683c8a99025fc28428849a29d5e01f4393f6860924024916df649c465ebee82c109132caf3e003263e22635a14088d1345afab411a612f7febbddd0223d4cf444fb344406f487ba3bb92c5db2dc8114013248be70fbfb9387f3a3b3d993d7af2f2e7af5f9dfdfafac5fd8bf3ef965717cf4e5f9d9dcfbe7f683002b03e1c1e82df072f2f1115d0e9bc0f7a48050200e0498f703339f2b67b44afd4b5ec7d13668fbf9afff648b2d86d7dfde2feecf9d38b670fe62767f31fff9a9dfeb4c431ffe3cbf9ef7fce1f3c972c36877ffff2f2c9b7f393b3d90f8f5daf9de820d93d0544f7e0f3c3b6ff6ee76d0f3a1dd8dab277db3235ca625784a7305054823f017cb07f036edffe0cc3148e7ae0f316f8030ded2a27f3677d580c589e2aaa332560c71ef419b21fabada2e36442412ada67c7108468294da6e875298d63919713299b38087181aad219c0b40f25527ca7ebe52b8dc59d2e864b16f6ff14c8d108aee55231a1c1db2daeb52c90311951483345410f8936e381a5404064e32e5501fa57015744a808007b97beaedfa29a344e1357d236b7e55d1b152bf94c3551bacc6625458bdcb6ea5837586597f802b047a404911c19b22a1382890110ae2889bf78c70d962cbe02a9488699046b4e01ee8199475514e0e58b1756c01e84319d8422e31c76f7b676600b01a49c5269479548cbe8790df8afa9b354bcc655521b7e56321a074150726a5418b58dea3ca5f5b63e619cc6a0133720d824c3aa338af284c4ffd79a11e31cfc4f3ff978773528f6d15caae6a60b7a74b0be6091316b4bb9279331d18e8b97bbb602d76295c83702bd9aee71f54770c164134683405a97eac888ced23792f8c5c63210d584d723b028b5a81a78dd6f4d18bf950c2ac8ba24f2b6bb24a5828c69f9360c3c536ba4b817f26490865ede2585fb063c1998b18e00f67630441160b317af52300788f2b532a274ad8e28bd5ea8e8a6d24456627565af0976a5d51d6dc4a3b3b40e4867e97aa991a45258ea894a234a1bda0ed5d4809d962ca66efdd48d9eda49a8cfd03f010000ffff68dc57ed42080000",
		"d2300211bf42dd2da9894b2bd073f15c": "1f8b08000000000000ff548ecb6aeb301086f77a0cc1594697f89026029182bb31b8c5a4982e8bac8c23515932d2a497b72f36bdd0cdc0cc7cffc77ff7f8a0694e09d5326e43b226b854505595943c2142a6a4f333041fa1f505fb53db99e51c357588b3e25c6e6f98608249b53d08c9e72ffa39f882cce174b4cec408a139eb7f8592fa7b2b9a4a4aea6bc134dd03666ffb1c7ead872d93bb3d934caa4a08b1e77625ff266a837049f943d3b5eb6600831b0726a0db5807f68592268ee1fafe943d427f6a7ffc03f3cc46258590fff9dbf23d9e073dadd64249977dc43ac5d15ff4684201d24484fc6a82a63b5128f90c0000ffff9990d67e3a010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("myBox", "assets")
		b.SetResolver("cnf.tpl.toml", packr.Pointer{ForwardBox: gk, ForwardPath: "d2300211bf42dd2da9894b2bd073f15c"})
		b.SetResolver("ctl.tpl.sh", packr.Pointer{ForwardBox: gk, ForwardPath: "cb62e23f86addb96ead8c397f3769a97"})
	}()
	return nil
}()
