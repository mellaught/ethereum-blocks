package eth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mellaught/ethereum-blocks/src/common"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// BlocksHandle ...
// Input:
// Output:
func (e *EthereumSRV) BlocksHandle(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)["input"]
	e.logger.WithField("Request", request).Infoln("BlocksHandle()")
	if strings.Contains(request, "-") {
		e.logger.Debugln("Try to find by range ", request)
		inputRange := strings.Split(strings.TrimSpace(request), "-")
		rangeLen := len(inputRange)
		if rangeLen != 2 {
			e.logger.WithFields(logrus.Fields{"function": "BlocksHandle()", "Range length": rangeLen}).
				Errorln("While check range length")
			common.ResponError(w, http.StatusBadRequest, fmt.Sprintf("Bad request range length %d", rangeLen))
			return
		}

		// try to convert input string range -> start, end block number(uint64)
		start, err := strconv.ParseUint(inputRange[0], 10, 64)
		if err != nil {
			e.logger.WithFields(logrus.Fields{"function": "BlocksHandle()", "Start": inputRange[0]}).
				Errorln("While parse uint", err)
			common.ResponError(w, http.StatusBadRequest, fmt.Sprintf("While parse range :%s", err.Error()))
			return
		}

		end, err := strconv.ParseUint(inputRange[1], 10, 64)
		if err != nil {
			e.logger.WithFields(logrus.Fields{"function": "BlocksHandle()", "End": inputRange[0]}).
				Errorln("While parse uint", err)
			common.ResponError(w, http.StatusBadRequest, fmt.Sprintf("While parse range :%s", err.Error()))
			return
		}

		blocks, err := e.blocks.GetBlocksByRange(start, end)
		if err != nil {
			e.logger.WithFields(logrus.Fields{"function": "BlocksHandle()", "End": inputRange[0]}).
				Errorln("While parse uint", err)
			common.ResponError(w, http.StatusBadRequest, fmt.Sprintf("While parse range :%s", err.Error()))
			return
		}
		common.ResponJSON(w, http.StatusOK, blocks)
		return
	}
	e.logger.Debugln("Try to find by transaction hash ", request)
	// check transaction id
	if _, err := hexutil.Decode(request); err != nil {
		e.logger.WithFields(logrus.Fields{"function": "BlocksHandle()", "Transaction hash": request}).
			Errorln("Bad transaction hash: ", err)
		common.ResponError(w, http.StatusBadRequest, "Bad transaction hash")
		return
	}
	blocks := e.blocks.GetBlockByTransactionID(request)

	common.ResponJSON(w, http.StatusOK, blocks)
}
