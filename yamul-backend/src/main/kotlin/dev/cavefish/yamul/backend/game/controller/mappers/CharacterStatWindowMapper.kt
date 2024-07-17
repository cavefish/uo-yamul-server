package dev.cavefish.yamul.backend.game.controller.mappers

import dev.cavefish.yamul.backend.game.api.MsgStatWindowInfo
import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import org.springframework.stereotype.Service

@Service
@SuppressWarnings("MaxLineLength", "MagicNumber")
class CharacterStatWindowMapper(
    private val objectIdMapper: ObjectIdMapper
) {

    fun map(gameObject: GameObject): MsgStatWindowInfo.Builder? {
        return MsgStatWindowInfo.newBuilder()
            .setCharacterID(objectIdMapper.create(gameObject.id))
            .setCharacterName(gameObject.name)
            .setHitPointsCurrent(45)
            .setHitPointsMax(45)
            .setFlagNameAllowed(false)
            .setLevel2(
                createLevel2()
            )
    }

    private fun createLevel2(): MsgStatWindowInfo.MsgStatWindowInfoLevel2.Builder? =
        MsgStatWindowInfo.MsgStatWindowInfoLevel2.newBuilder()
            .setStrength(0)
            .setStrength(45)
            .setIntelligence(35)
            .setStaminaCurrent(10)
            .setStaminaMax(35)
            .setManaCurrent(35)
            .setManaMax(10)
            .setGold(655360)
            .setResistancePhysical(1000)
            .setWeightCurrent(0)
            .setLevel3(
                createLevel3()
            )

    private fun createLevel3(): MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.Builder? =
        MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.newBuilder()
            .setStatsCap(50433)
            .setLevel4(
                createLevel4()
            )

    private fun createLevel4(): MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.Builder? =
        MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.newBuilder()
            .setFollowersCurrent(1)
            .setFollowersMax(44)
            .setLevel5(
                createLevel5()
            )

    private fun createLevel5(): MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.MsgStatWindowInfoLevel5.Builder? =
        MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.MsgStatWindowInfoLevel5.newBuilder()
            .setResistanceFire(5)
            .setResistanceCold(0)
            .setResistancePoison(0)
            .setResistanceEnergy(0)
            .setLuck(0)
            .setDamageMin(0)
            .setDamageMax(1)
            .setTithingPoints(262144)
            .setLevel6(
                createLevel6()
            )

    private fun createLevel6(): MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.MsgStatWindowInfoLevel5.MsgStatWindowInfoLevel6.Builder? =
        MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.MsgStatWindowInfoLevel5.MsgStatWindowInfoLevel6.newBuilder()
            .setWeightMax(72)
            .setRace(0)
            .setLevel7(
                createLevel7()
            )

    private fun createLevel7(): MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.MsgStatWindowInfoLevel5.MsgStatWindowInfoLevel6.MsgStatWindowInfoLevel7.Builder? =
        MsgStatWindowInfo.MsgStatWindowInfoLevel2.MsgStatWindowInfoLevel3.MsgStatWindowInfoLevel4.MsgStatWindowInfoLevel5.MsgStatWindowInfoLevel6.MsgStatWindowInfoLevel7.newBuilder()
            .setResistancePhysicalMax(0)
            .setResistanceFireMax(70)
            .setResistanceColdMax(70)
            .setResistancePoisonMax(70)
            .setResistanceEnergyMax(70)
            .setDefenseChanceIncreaseCurrent(70)
            .setDefenseChanceIncreaseMax(0)
            .setHitChanceIncrease(45)
            .setSwingSpeedIncrease(0)
            .setDamageIncrease(0)
            .setLowerReagentCost(0)
            .setSpellDamageIncrease(0)
            .setFasterCasting(0)
            .setFasterCastRecovery(0)
            .setLowerManaCost(0)
}