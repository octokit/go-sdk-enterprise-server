package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207 "github.com/octokit/go-sdk-enterprise-server/pkg/github/models"
)

type ItemItemRulesetsItemWithRuleset_PutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The actors that can bypass the rules in this ruleset
    bypass_actors []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetBypassActorable
    // Parameters for a repository ruleset ref name condition
    conditions ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetConditionsable
    // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page.
    enforcement *ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleEnforcement
    // The name of the ruleset.
    name *string
    // An array of rules within the ruleset.
    rules []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleable
}
// NewItemItemRulesetsItemWithRuleset_PutRequestBody instantiates a new ItemItemRulesetsItemWithRuleset_PutRequestBody and sets the default values.
func NewItemItemRulesetsItemWithRuleset_PutRequestBody()(*ItemItemRulesetsItemWithRuleset_PutRequestBody) {
    m := &ItemItemRulesetsItemWithRuleset_PutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemRulesetsItemWithRuleset_PutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemRulesetsItemWithRuleset_PutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemRulesetsItemWithRuleset_PutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBypassActors gets the bypass_actors property value. The actors that can bypass the rules in this ruleset
// returns a []RepositoryRulesetBypassActorable when successful
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) GetBypassActors()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetBypassActorable) {
    return m.bypass_actors
}
// GetConditions gets the conditions property value. Parameters for a repository ruleset ref name condition
// returns a RepositoryRulesetConditionsable when successful
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) GetConditions()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetConditionsable) {
    return m.conditions
}
// GetEnforcement gets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page.
// returns a *RepositoryRuleEnforcement when successful
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) GetEnforcement()(*ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleEnforcement) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bypass_actors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRepositoryRulesetBypassActorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetBypassActorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetBypassActorable)
                }
            }
            m.SetBypassActors(res)
        }
        return nil
    }
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRepositoryRulesetConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetConditionsable))
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.ParseRepositoryRuleEnforcement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val.(*ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleEnforcement))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.CreateRepositoryRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleable)
                }
            }
            m.SetRules(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the ruleset.
// returns a *string when successful
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) GetName()(*string) {
    return m.name
}
// GetRules gets the rules property value. An array of rules within the ruleset.
// returns a []RepositoryRuleable when successful
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) GetRules()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleable) {
    return m.rules
}
// Serialize serializes information the current object
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBypassActors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBypassActors()))
        for i, v := range m.GetBypassActors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("bypass_actors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcement() != nil {
        cast := (*m.GetEnforcement()).String()
        err := writer.WriteStringValue("enforcement", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBypassActors sets the bypass_actors property value. The actors that can bypass the rules in this ruleset
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) SetBypassActors(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetBypassActorable)() {
    m.bypass_actors = value
}
// SetConditions sets the conditions property value. Parameters for a repository ruleset ref name condition
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) SetConditions(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetConditionsable)() {
    m.conditions = value
}
// SetEnforcement sets the enforcement property value. The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page.
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) SetEnforcement(value *ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleEnforcement)() {
    m.enforcement = value
}
// SetName sets the name property value. The name of the ruleset.
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) SetName(value *string)() {
    m.name = value
}
// SetRules sets the rules property value. An array of rules within the ruleset.
func (m *ItemItemRulesetsItemWithRuleset_PutRequestBody) SetRules(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleable)() {
    m.rules = value
}
type ItemItemRulesetsItemWithRuleset_PutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBypassActors()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetBypassActorable)
    GetConditions()(ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetConditionsable)
    GetEnforcement()(*ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleEnforcement)
    GetName()(*string)
    GetRules()([]ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleable)
    SetBypassActors(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetBypassActorable)()
    SetConditions(value ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRulesetConditionsable)()
    SetEnforcement(value *ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleEnforcement)()
    SetName(value *string)()
    SetRules(value []ie1e2072a5a4eb80f74a1387d59644d3f70804e6b7b2f406016da8826571f1207.RepositoryRuleable)()
}
