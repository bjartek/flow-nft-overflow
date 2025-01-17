package main

import (
	"fmt"

	. "github.com/bjartek/overflow"
	"github.com/fatih/color"
)

func main() {

	fmt.Print("\033[H\033[2J")
	color.Green("This is a demo of what overflow can do for an NFT project, in this case the example is a modified version of flow-nft repo")

	color.Green("In order to start overflow with a default in memory client you simply run")
	fmt.Println("")
	color.Cyan(`o := Overflow()`)
	color.Green("")
	color.Green("This will start overflow in embedded mode, if any interaction error occur the job will terminate in embedded mode")
	pause()
	o := Overflow()
	fmt.Println("")
	color.Green(`We now have an running version of 'Overflow' in embedded mode, ready to be interacted with. All stakeholders are created with the default amount of flow and contracts are deployed.`)
	pause()

	color.Green("Running an transaction in overflow is done by calling the `Tx` method on the `o` or overflow object")
	color.Green(`We now have an running version of 'Overflow' in embedded mode, ready to be interacted with`)
	fmt.Println("")
	color.Cyan(`o.Tx("setup_account", WithSigner("alice"))`)
	color.Green("")
	color.Green("This will run the transaction `setup_account` from the `transactions/` folder, sign is as the demo user `alice` it will also print out the result in a nice terse way")
	color.Green("")
	color.Green("Note that when we refer to users by name in overflow we do not use the network prefix, this is so that you can have the same stakeholders on mainnet/testnet if you want to without chaning the code. So in flow.json the account for alice is called 'emulator-alice'")

	pause()
	o.Tx("setup_account", WithSigner("alice"))

	color.Green("Now we have set up alice, lets set up bob and also setup their royalty receivers")

	color.Green("")
	color.Cyan(`
	o.Tx("setup_account", 
		WithSigner("bob")
	)

	o.Tx("setup_account_to_receive_royalty", 
		WithSigner("alice"), 
		WithArg("vaultPath", "/storage/flowTokenVault"),
	)
	o.Tx("setup_account_to_receive_royalty", 
		WithSigner("bob"), 
		WithArg("vaultPath", "/storage/flowTokenVault"),
	)	
	`)

	color.Green("")
	o.Tx("setup_account",
		WithSigner("bob"),
	)
	o.Tx("setup_account_to_receive_royalty",
		WithSigner("alice"),
		WithArg("vaultPath", "/storage/flowTokenVault"),
	)
	o.Tx("setup_account_to_receive_royalty",
		WithSigner("bob"),
		WithArg("vaultPath", "/storage/flowTokenVault"),
	)

	color.Green("Everything is now ready to mint an NFT into alice collection!")

	pause()

	color.Green(`Minting is running another transaction but in thise case we have a lot more arguments to the transaction.

In overflow v1 all arguments are _named_ that is you mention them by their name and the value and not the order they appear in the transaction. If you use the wrong names and types then overflow will let you know with an terse appropritate error message`)

	color.Cyan(`
  id,_ :=o.Tx("mint_nft", 
	  WithSignerServiceAccount(),
		WithArg("recipient", "alice"),
		WithArg("name", "Example NFT 0"),
		WithArg("description", "This is an example NFT"),
		WithArg("thumbnail", "example.jpeg"),
		WithArg("cuts", "[0.25, 0.40]"),
		WithArg("royaltyDescriptions", ` + "`" + `["minter","creator"]` + "`" + `'),
		WithAddresses("royaltyBeneficiaries", "alice", "bob")).
		GetIdFromEvent("Deposit", "id")
		`)

	color.Green("Most arguments in overflow are sent using the `Arg` method but there are some other helpfull methods, in this case we use `Addresses` to return a list of addresses. As you can see we can use the logical name of the account in flow.json and it will replace that with the address in the transaction")

	color.Green("We can also see that after we have run and printed the result we can fetch out data from the events in the transaction, in this case we fetch out the first entry of an event that has the suffix Deposit and we fetch the id `id` field as an UInt64. This is a convenience method that was added since this is a very normal pattern")

	pause()

	id, _ := o.Tx("mint_nft", WithSignerServiceAccount(),
		WithArg("recipient", "alice"),
		WithArg("name", "Example NFT 0"),
		WithArg("description", "This is an example NFT"),
		WithArg("thumbnail", "example.jpeg"),
		WithArg("cuts", "[0.25, 0.40]"),
		WithArg("royaltyDescriptions", `["minter","creator"]`),
		WithAddresses("royaltyBeneficiaries", "alice", "bob")).
		GetIdFromEvent("Deposit", "id")

	color.Green("We now have an NFT that is minted with id %d that we can run some scripts against!\n", id)

	pause()

	color.Green("A script is run in very much the same way as a Transaction only it uses the `Script` method like the following example")
	color.Cyan(`o.Script("get_nft_metadata", WithArg("address", "alice"), WithArg("id", id))`)

	pause()
	o.Script("get_nft_metadata", WithArg("address", "alice"), WithArg("id", id))

	color.Green("And that is the metadata of the nft that we just minted, hope you like what overflow can do to tell a story! Oh and if you want to run this story against `testnet` you can easily do that. At .find we use overflow to run system.d job, cronjobs, serverless functions and lots of things. it is the green goo that keeps everything (over)flowing")

}

func pause() {
	fmt.Println()
	color.Yellow("press any key to continue")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")
}
