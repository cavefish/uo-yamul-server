package dev.cavefish.yamul

import com.flextrade.jfixture.JFixture
import org.assertj.core.api.SoftAssertions
import org.assertj.core.api.junit.jupiter.InjectSoftAssertions
import org.assertj.core.api.junit.jupiter.SoftAssertionsExtension
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.extension.ExtendWith
import org.springframework.boot.SpringBootConfiguration
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.ComponentScan
import org.springframework.test.context.ActiveProfiles

@ComponentScan(basePackages = ["dev.cavefish.yamul.backend"])
@SpringBootTest
@SpringBootConfiguration
@ActiveProfiles("test")
@ExtendWith(SoftAssertionsExtension::class)
open class IntegrationTest {
    @InjectSoftAssertions
    protected lateinit var softly: SoftAssertions

    protected var fixture: JFixture = JFixture()

    @BeforeEach
    fun beforeEach() {
        fixture = JFixture()
    }

}