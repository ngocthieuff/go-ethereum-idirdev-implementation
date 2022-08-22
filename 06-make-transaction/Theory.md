:sparkles: An ethereum transaction is composed of following components :sparkles:
{
 `nonce` : 'how many confirmed transactions this account has sent previously?',
 `gasPrice` : 'price of gas (in wei) the originator is willing to pay for this transaction',
 `to` : 'recipient of this transaction(either smart contract or EOA)',
 `value` : 'how much ether this transaction pays',
 `data` : 'any binary data payload',
 `v,r,s` : 'three components of ECDSA signature of originating account holder'
}