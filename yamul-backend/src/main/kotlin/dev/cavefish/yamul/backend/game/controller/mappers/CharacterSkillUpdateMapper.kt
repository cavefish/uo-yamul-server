package dev.cavefish.yamul.backend.game.controller.mappers

import dev.cavefish.yamul.backend.game.api.MsgSkillUpdateServer
import dev.cavefish.yamul.backend.game.api.MsgSkillUpdateStatus
import dev.cavefish.yamul.backend.game.api.MsgSkillUpdateType
import org.springframework.stereotype.Service

@Service
@SuppressWarnings("MagicNumber")
class CharacterSkillUpdateMapper {

    fun getFullUpdate():MsgSkillUpdateServer.Builder {
        return MsgSkillUpdateServer
            .newBuilder()
            .setType(MsgSkillUpdateType.basicSkillCap)
            .addAllSkills(IntRange(1, 46).map { getSkillItem(it).build() })
    }

    private fun getSkillItem(id: Int) = MsgSkillUpdateServer.MsgSkillUpdateSkills
        .newBuilder()
        .setSkillId(id)
        .setValue(300)
        .setMaxValue(1000)
        .setBaseValue(250)
        .setStatus(MsgSkillUpdateStatus.locked)

}