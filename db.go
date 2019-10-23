package rossby

import (
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/kansaslabs/rossby/bucket"
	pb "github.com/kansaslabs/rossby/pb"
)

// Find a BoxID from a list of contacts. Returns an error if the box is not found or if
// there are more than one boxes associated with the specified list of contacts.
func (r *Replica) boxFromContacts(contacts []*pb.Contact) (box BoxID, err error) {
	if err = r.db.View(func(txn *badger.Txn) error {
		boxset := make(map[BoxID]struct{})

		for _, contact := range contacts {
			key := bucket.Contacts.Key([]byte(contact.Contact))
			item, err := txn.Get(key)
			if err == badger.ErrKeyNotFound {
				continue
			} else if err != nil {
				return err
			}

			idb, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			id, err := ParseBoxID(idb)
			if err != nil {
				return err
			}
			boxset[id] = struct{}{}
		}

		if len(boxset) == 0 {
			return fmt.Errorf("no box id found for specified contacts")
		}

		if len(boxset) > 1 {
			return fmt.Errorf("multiple box ids discovered for contacts list")
		}

		// Fetch the boxid from the boxset
		for b := range boxset {
			box = b
			break
		}

		return nil
	}); err != nil {
		return NilBox, err
	}

	return box, nil
}
