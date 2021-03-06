/*
Copyright (C) 2018 Expedia Group.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package command

import (
	"encoding/json"
	"fmt"
	"github.com/HotelsDotCom/flyte-bamboo/event"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/go-logger"
)

func (c CommandService) EnablePlanCommand() flyte.Command {
	return flyte.Command{
		Name:         "EnablePlan",
		OutputEvents: []flyte.EventDef{event.EnablePlanSuccessEventDef, event.EnablePlanErrorEventDef},
		Handler:      c.EnablePlanHandler,
	}
}

func (c CommandService) EnablePlanHandler(input json.RawMessage) flyte.Event {

	var handlerInput struct {
		Plan string `json:"plan"`
	}

	if err := json.Unmarshal(input, &handlerInput); err != nil {
		err = fmt.Errorf("could not marshal EnablePlan command input: %v", err)
		logger.Error(err)
		return flyte.NewFatalEvent(fmt.Sprintf("Input is invalid: %v", err))
	}

	if err := c.bambooClient.EnablePlan(handlerInput.Plan); err != nil {
		err = fmt.Errorf("could not enable plan: %v\n", err)
		logger.Error(err)
		return event.EnablePlanErrorEvent(fmt.Sprintf("Fail: %s", err), handlerInput.Plan)
	}

	return event.EnablePlanSuccessEvent(handlerInput.Plan)

}
